package authv3

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/twin-te/twinte-back/appenv"
	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
	"golang.org/x/oauth2"
)

var appleOAuth2Config = &oauth2.Config{
	ClientID:     appenv.OAUTH2_APPLE_CLIENT_ID,
	ClientSecret: appenv.OAUTH2_APPLE_CLIENT_SECRET,
	Endpoint: oauth2.Endpoint{
		AuthURL:   "https://appleid.apple.com/auth/oauth2/v2/authorize",
		TokenURL:  "https://appleid.apple.com/auth/oauth2/v2/token",
		AuthStyle: oauth2.AuthStyleInParams,
	},
	RedirectURL: appenv.OAUTH2_APPLE_CALLBACK_URL,
	Scopes:      []string{""},
}

func getAppleSocialID(ctx context.Context, code string) (socialID authdomain.SocialID, err error) {
	token, err := appleOAuth2Config.Exchange(ctx, code)
	if err != nil {
		return
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		return "", errors.New("failed to retrieve id token")
	}

	verificationKeySet, err := getApplePublicKeySet(ctx)
	if err != nil {
		return
	}

	t, err := jwt.Parse(
		idToken,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return verificationKeySet, nil
		},
		jwt.WithAudience(appleOAuth2Config.ClientID),
		jwt.WithIssuer("https://appleid.apple.com"),
	)
	if err != nil {
		return
	}

	// TODO: verify id token
	// c.f. https://developer.apple.com/documentation/sign_in_with_apple/sign_in_with_apple_rest_api/verifying_a_user#3383769

	sub, err := t.Claims.GetSubject()
	if err != nil {
		return
	}

	return authdomain.ParseSocialID(sub)
}

func getApplePublicKeySet(ctx context.Context) (verificationKeySet jwt.VerificationKeySet, err error) {
	// c.f. https://developer.apple.com/documentation/sign_in_with_apple/fetch_apple_s_public_key_for_verifying_token_signature
	resp, err := http.Get("https://appleid.apple.com/auth/keys")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	v := &struct {
		Keys []struct {
			Alg string `json:"alg"`
			E   string `json:"e"`
			Kid string `json:"kid"`
			Kty string `json:"kty"`
			N   string `json:"n"`
			Use string `json:"use"`
		} `json:"keys"`
	}{}

	err = json.Unmarshal(body, v)

	for _, key := range v.Keys {
		publicKey := &rsa.PublicKey{
			E: int(decodeBase64BigInt(key.E).Int64()),
			N: decodeBase64BigInt(key.E),
		}
		verificationKeySet.Keys = append(verificationKeySet.Keys, publicKey)
	}

	return
}

// decodeBase64BigInt decodes a base64-encoded larger integer from Apple's key format.
// c.f. https://stackoverflow.com/questions/66067321/marshal-appleids-public-key-to-rsa-publickey
func decodeBase64BigInt(s string) *big.Int {
	buffer, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(s)
	if err != nil {
		log.Fatalf("failed to decode base64: %v", err)
	}
	return big.NewInt(0).SetBytes(buffer)
}
