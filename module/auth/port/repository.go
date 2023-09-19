package authport

import (
	"context"
	"time"

	"github.com/twin-te/twinte-back/idtype"
	authentity "github.com/twin-te/twinte-back/module/auth/entity"
)

type Repository interface {
	Transaction(ctx context.Context, fc func(rtx Repository) error, readOnly bool) error

	SaveUser(ctx context.Context, user *authentity.User) error
	FindUser(ctx context.Context, conds FindUserConds) (*authentity.User, error)
	DeleteUser(ctx context.Context, id idtype.UserID) error

	SaveSession(ctx context.Context, session *authentity.Session) error
	FindSession(ctx context.Context, conds FindSessionConds) (*authentity.Session, error)
	DeleteSessions(ctx context.Context, conds DeleteSessionsConds) error
}

type FindUserConds struct {
	ID                 *idtype.UserID
	UserAuthentication *authentity.UserAuthentication
}

type FindSessionConds struct {
	ID             *idtype.SessionID
	UserID         *idtype.UserID
	ExpiredAtAfter *time.Time
}

type DeleteSessionsConds struct {
	IDs     *idtype.SessionIDs
	UserIDs *idtype.UserIDs
}
