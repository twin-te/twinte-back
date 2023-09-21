package authgateway

import (
	"context"
	_ "embed"

	"github.com/twin-te/twinte-back/appenv"
	authentity "github.com/twin-te/twinte-back/module/auth/entity"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleapioauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var googleOAuth2Config = &oauth2.Config{
	ClientID:     appenv.OAUTH2_GOOGLE_CLIENT_ID,
	ClientSecret: appenv.OAUTH2_GOOGLE_CLIENT_SECRET,
	Endpoint:     google.Endpoint,
	RedirectURL:  appenv.OAUTH2_GOOGLE_CALLBACK_URL,
	Scopes:       []string{googleapioauth2.OpenIDScope},
}

func (uc *Impl) GetGoogleOAuth2ConsentPageURL(ctx context.Context, state authentity.OAuth2State) (url authentity.OAuth2ConsentPageURL, err error) {
	url = authentity.NewOAuth2ConsentPageURLFromString(googleOAuth2Config.AuthCodeURL(state.String(), oauth2.ApprovalForce))
	return
}

func (uc *Impl) GetGoogleSocialIDFromCode(ctx context.Context, code authentity.OAuth2Code) (authentity.SocialID, error) {
	token, err := googleOAuth2Config.Exchange(ctx, code.String())
	if err != nil {
		return "", err
	}

	client := googleOAuth2Config.Client(ctx, token)

	srv, err := googleapioauth2.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return "", err
	}

	tokenInfo, err := srv.Tokeninfo().Do()
	if err != nil {
		return "", err
	}

	return authentity.NewSocialIDFromString(tokenInfo.UserId), nil
}
