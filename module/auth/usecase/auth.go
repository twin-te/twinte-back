package authusecase

import (
	"context"
	"errors"
	"time"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/appctx"
	"github.com/twin-te/twinte-back/apperr"
	"github.com/twin-te/twinte-back/idtype"
	authentity "github.com/twin-te/twinte-back/module/auth/entity"
	authport "github.com/twin-te/twinte-back/module/auth/port"
)

func (uc *Impl) Authenticate(ctx context.Context, id idtype.SessionID) (*authentity.User, error) {
	session, err := uc.r.FindSession(ctx, authport.FindSessionConds{
		ID:             &id,
		ExpiredAtAfter: lo.ToPtr(time.Now()),
	})
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return nil, apperr.ErrUnauthenticated
		}
		return nil, err
	}

	user, err := uc.r.FindUser(ctx, authport.FindUserConds{
		ID: &session.UserID,
	})
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return nil, apperr.ErrUnauthenticated
		}
		return nil, err
	}

	return user, nil
}

func (uc *Impl) AuthorizePublic(ctx context.Context) error {
	return nil
}

func (uc *Impl) AuthorizeAuthenticatedUser(ctx context.Context) (*authentity.User, error) {
	user, ok := appctx.GetUser(ctx)
	if !ok {
		return nil, apperr.ErrUnauthenticated
	}
	return user, nil
}
