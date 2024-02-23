package authmodule

import (
	"context"

	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

type AccessController interface {
	WithActor(ctx context.Context, id *idtype.SessionID) (context.Context, error)

	// Authenticate returns user id if the actor is authenticated.
	//
	// [Error Code]
	//   - shared.Unauthenticated
	Authenticate(ctx context.Context) (idtype.UserID, error)

	// Authorize verifies that the actor has the given permission.
	//
	// [Error Code]
	//   - shared.Unauthorized
	Authorize(ctx context.Context, permission authdomain.Permission) error
}

// UseCase represents application specific business rules.
//
// The error codes for authentication and authorization failures are not stated explicitly.
type UseCase interface {
	// SignUpOrLogin creates a new user, if the user identified by the given authentication information does not exist.
	// Next, a new session will be created, if there is no valid session associated with the user identified by the given authentication information.
	//
	// [Authentication] not required
	SignUpOrLogin(ctx context.Context, userAuthentication authdomain.UserAuthentication) (*authdomain.Session, error)

	// GetMe returns the user.
	//
	// [Authentication] required
	GetMe(ctx context.Context) (*authdomain.User, error)

	// AddUserAuthentication adds the given authentication to the user.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - shared.AlreadyExists ( if the given authentication information already exists )
	//   - auth.MultipleAuthenticationOfSameProvider
	AddUserAuthentication(ctx context.Context, userAuthentication authdomain.UserAuthentication) error

	// DeleteUserAuthentication removes the authentication information of the given provider from the user.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - shared.NotFound ( if the authentication specified by the given provider is not found )
	//   - auth.UserHasAtLeastOneAuthentication
	DeleteUserAuthentication(ctx context.Context, provider authdomain.Provider) error

	// Logout deletes all sessions associated with the user.
	//
	// [Authentication] required
	Logout(ctx context.Context) error

	// DeleteAccount deletes the account and all sessions associated with the user.
	//
	// [Authentication] required
	DeleteAccount(ctx context.Context) error
}
