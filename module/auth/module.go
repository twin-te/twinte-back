package authmodule

import (
	"context"

	"github.com/twin-te/twinte-back/idtype"
	authentity "github.com/twin-te/twinte-back/module/auth/entity"
)

type UseCase interface {
	GenerateOAuth2State(ctx context.Context) (authentity.OAuth2State, error)
	GetOAuth2ConsentPageURL(ctx context.Context, provider authentity.Provider, state authentity.OAuth2State) (authentity.OAuth2ConsentPageURL, error)
	GetSocialIDFromCode(ctx context.Context, provider authentity.Provider, code authentity.OAuth2Code) (authentity.SocialID, error)

	Authenticate(ctx context.Context, id idtype.SessionID) (*authentity.User, error)

	AuthorizePublic(ctx context.Context) error
	AuthorizeAuthenticatedUser(ctx context.Context) (*authentity.User, error)

	SignUpOrLogin(ctx context.Context, authentication authentity.UserAuthentication) (*authentity.Session, error)
	Logout(ctx context.Context, user *authentity.User) error
	DeleteAccount(ctx context.Context, user *authentity.User) error

	AddUserAuthentication(ctx context.Context, user *authentity.User, authentication authentity.UserAuthentication) error
	DeleteUserAuthentication(ctx context.Context, user *authentity.User, provider authentity.Provider) error
}
