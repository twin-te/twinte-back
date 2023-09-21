package appctx

import (
	"context"

	authentity "github.com/twin-te/twinte-back/module/auth/entity"
)

type userKey struct{}

func GetUser(ctx context.Context) (*authentity.User, bool) {
	user, ok := ctx.Value(userKey{}).(*authentity.User)
	return user, ok
}

func SetUser(ctx context.Context, user *authentity.User) context.Context {
	return context.WithValue(ctx, userKey{}, user)
}
