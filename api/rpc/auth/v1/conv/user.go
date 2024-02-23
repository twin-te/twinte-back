package authv1conv

import (
	sharedconv "github.com/twin-te/twinte-back/api/rpc/shared/conv"
	authv1 "github.com/twin-te/twinte-back/api/rpcgen/auth/v1"
	"github.com/twin-te/twinte-back/base"
	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
)

func ToPBUser(user *authdomain.User) (*authv1.User, error) {
	pbAuthentications, err := base.MapWithErr(user.Authentications, ToPBUserAuthentication)
	if err != nil {
		return nil, err
	}

	pbUser := &authv1.User{
		Id:              sharedconv.ToPBUUID(user.ID),
		Authentications: pbAuthentications,
		CreatedAt:       sharedconv.ToPBRFC3339DateTime(user.CreatedAt),
	}

	return pbUser, nil
}
