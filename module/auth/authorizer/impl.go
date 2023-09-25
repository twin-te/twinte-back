package authorizer

import (
	"context"

	"github.com/twin-te/twinte-back/appctx"
	"github.com/twin-te/twinte-back/apperr"
	authentity "github.com/twin-te/twinte-back/module/auth/entity"
)

type impl struct{}

func (uc *impl) AuthorizePublic(ctx context.Context) error {
	return nil
}

func (uc *impl) AuthorizeAuthenticatedUser(ctx context.Context) (*authentity.User, error) {
	user, ok := appctx.GetUser(ctx)
	if !ok {
		return nil, apperr.ErrUnauthenticated
	}
	return user, nil
}

func New() *impl {
	return &impl{}
}
