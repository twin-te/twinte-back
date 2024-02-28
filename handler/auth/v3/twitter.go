package authv3

import (
	"context"
	"encoding/json"
	"io"

	"github.com/twin-te/twinte-back/appenv"
	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
	"golang.org/x/oauth2"
)

var twitterOAuth2Config = &oauth2.Config{
	ClientID:     appenv.OAUTH2_TWITTER_CLIENT_ID,
	ClientSecret: appenv.OAUTH2_TWITTER_CLIENT_SECRET,
	Endpoint: oauth2.Endpoint{
		AuthURL:   "https://twitter.com/i/oauth2/authorize",
		TokenURL:  "https://api.twitter.com/2/oauth2/token",
		AuthStyle: oauth2.AuthStyleInHeader,
	},
	RedirectURL: appenv.OAUTH2_TWITTER_CALLBACK_URL,
	Scopes:      []string{"users.read", "tweet.read"},
}

func getTwitterSocialID(ctx context.Context, code string) (socialID authdomain.SocialID, err error) {
	token, err := twitterOAuth2Config.Exchange(ctx, code, verifierOption)
	if err != nil {
		return
	}

	client := twitterOAuth2Config.Client(ctx, token)
	resp, err := client.Get("https://api.twitter.com/2/users/me")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	v := &struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}{}

	if err = json.Unmarshal(body, v); err != nil {
		return
	}

	return authdomain.ParseSocialID(v.Data.ID)
}
