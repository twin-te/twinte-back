package authusecase

import (
	"context"
	"errors"

	"github.com/twin-te/twinte-back/apperr"
	"github.com/twin-te/twinte-back/idtype"
	authentity "github.com/twin-te/twinte-back/module/auth/entity"
	authport "github.com/twin-te/twinte-back/module/auth/port"
)

func (uc Impl) SignUpOrLogin(ctx context.Context, authentication authentity.UserAuthentication) (*authentity.Session, error) {
	var user *authentity.User

	err := uc.r.Transaction(ctx, func(rtx authport.Repository) (err error) {
		user, err = rtx.FindUser(ctx, authport.FindUserConds{
			UserAuthentication: &authentication,
		})

		if errors.Is(err, apperr.ErrNotFound) {
			user = &authentity.User{
				ID:              idtype.NewUserID(),
				Authentications: []authentity.UserAuthentication{authentication},
			}
			return rtx.SaveUser(ctx, user)
		}
		return
	}, false)

	if err != nil {
		return nil, err
	}

	session := authentity.NewSession(user.ID)
	return session, uc.r.SaveSession(ctx, session)
}

func (uc Impl) Logout(ctx context.Context, user *authentity.User) error {
	return uc.r.DeleteSessions(ctx, authport.DeleteSessionsConds{
		UserIDs: &idtype.UserIDs{user.ID},
	})
}

func (uc Impl) DeleteAccount(ctx context.Context, user *authentity.User) error {
	return uc.r.DeleteUser(ctx, user.ID)
}

func (uc Impl) AddUserAuthentication(ctx context.Context, user *authentity.User, authentication authentity.UserAuthentication) error {
	if err := user.AddAuthentication(authentication); err != nil {
		return err
	}
	return uc.r.SaveUser(ctx, user)
}

func (uc Impl) DeleteUserAuthentication(ctx context.Context, user *authentity.User, provider authentity.Provider) error {
	if err := user.DeleteAuthentication(provider); err != nil {
		return err
	}

	return uc.r.SaveUser(ctx, user)
}
