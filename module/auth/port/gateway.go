package authport

import (
	"context"

	authentity "github.com/twin-te/twinte-back/module/auth/entity"
)

type Gateway interface {
	// Google
	GetGoogleOAuth2ConsentPageURL(ctx context.Context, state authentity.OAuth2State) (authentity.OAuth2ConsentPageURL, error)
	GetGoogleSocialIDFromCode(ctx context.Context, code authentity.OAuth2Code) (authentity.SocialID, error)
}
