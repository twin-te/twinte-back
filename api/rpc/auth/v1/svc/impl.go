package authv1svc

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"

	authv1conv "github.com/twin-te/twinte-back/api/rpc/auth/v1/conv"
	authv1 "github.com/twin-te/twinte-back/api/rpcgen/auth/v1"
	"github.com/twin-te/twinte-back/api/rpcgen/auth/v1/authv1connect"

	"github.com/twin-te/twinte-back/appenv"
	authmodule "github.com/twin-te/twinte-back/module/auth"
)

var _ authv1connect.AuthServiceHandler = (*impl)(nil)

type impl struct {
	uc authmodule.UseCase
}

func (svc *impl) GetMe(ctx context.Context, req *connect.Request[authv1.GetMeRequest]) (res *connect.Response[authv1.GetMeResponse], err error) {
	user, err := svc.uc.GetMe(ctx)
	if err != nil {
		return
	}

	pbUser, err := authv1conv.ToPBUser(user)
	if err != nil {
		return
	}

	res = connect.NewResponse(&authv1.GetMeResponse{
		User: pbUser,
	})

	return
}

func (svc *impl) Logout(ctx context.Context, req *connect.Request[authv1.LogoutRequest]) (res *connect.Response[authv1.LogoutResponse], err error) {
	if err = svc.uc.Logout(ctx); err != nil {
		return
	}

	res = connect.NewResponse(&authv1.LogoutResponse{})
	cookie := http.Cookie{
		Name:     appenv.COOKIE_SESSION_NAME,
		MaxAge:   -1,
		Secure:   appenv.COOKIE_SECURE,
		HttpOnly: true,
	}
	res.Header().Set("Set-Cookie", cookie.String())

	return
}

func (svc *impl) DeleteAccount(ctx context.Context, req *connect.Request[authv1.DeleteAccountRequest]) (res *connect.Response[authv1.DeleteAccountResponse], err error) {
	if err = svc.uc.DeleteAccount(ctx); err != nil {
		return
	}

	res = connect.NewResponse(&authv1.DeleteAccountResponse{})
	cookie := http.Cookie{
		Name:     appenv.COOKIE_SESSION_NAME,
		MaxAge:   -1,
		Secure:   appenv.COOKIE_SECURE,
		HttpOnly: true,
	}
	res.Header().Set("Set-Cookie", cookie.String())

	return
}

func (svc *impl) DeleteUserAuthentication(ctx context.Context, req *connect.Request[authv1.DeleteUserAuthenticationRequest]) (res *connect.Response[authv1.DeleteUserAuthenticationResponse], err error) {
	provider, err := authv1conv.FromPBProvider(req.Msg.Provider)
	if err != nil {
		return
	}

	if err = svc.uc.DeleteUserAuthentication(ctx, provider); err != nil {
		return
	}

	res = connect.NewResponse(&authv1.DeleteUserAuthenticationResponse{})

	return
}

func New(uc authmodule.UseCase) *impl {
	return &impl{uc: uc}
}
