package authusecase

import (
	"context"

	"fmt"

	"github.com/google/uuid"
	authentity "github.com/twin-te/twinte-back/module/auth/entity"
)

func (uc *Impl) GenerateOAuth2State(ctx context.Context) (authentity.OAuth2State, error) {
	return authentity.NewOAuth2StateFromString(uuid.NewString()), nil
}

func (uc *Impl) GetOAuth2ConsentPageURL(ctx context.Context, provider authentity.Provider, state authentity.OAuth2State) (authentity.OAuth2ConsentPageURL, error) {
	switch provider {
	case authentity.ProviderGoogle:
		return uc.g.GetGoogleOAuth2ConsentPageURL(ctx, state)
	}
	return "", fmt.Errorf("provider %d does not supported", provider)
}

func (uc *Impl) GetSocialIDFromCode(ctx context.Context, provider authentity.Provider, code authentity.OAuth2Code) (authentity.SocialID, error) {
	switch provider {
	case authentity.ProviderGoogle:
		return uc.g.GetGoogleSocialIDFromCode(ctx, code)
	}
	return "", fmt.Errorf("provider %d does not supported", provider)
}
