package restv3

import (
	"context"

	"github.com/twin-te/twinte-back/api/rest/v3/openapi"
	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
)

func toApiUser(user *authdomain.User) openapi.User {
	return openapi.User{
		Id: toApiUUID(user.ID),
	}
}

// ログイン中のユーザーを退会する
// (DELETE /users/me)
func (h *impl) DeleteUsersMe(ctx context.Context, request openapi.DeleteUsersMeRequestObject) (res openapi.DeleteUsersMeResponseObject, err error) {
	err = h.authUseCase.DeleteAccount(ctx)

	res = openapi.DeleteUsersMe204Response{}

	return
}

// ログイン中のユーザーを取得する
// (GET /users/me)
func (h *impl) GetUsersMe(ctx context.Context, request openapi.GetUsersMeRequestObject) (res openapi.GetUsersMeResponseObject, err error) {
	user, err := h.authUseCase.GetMe(ctx)
	if err != nil {
		return
	}

	apiUser := toApiUser(user)

	res = openapi.GetUsersMe200JSONResponse(apiUser)

	return
}
