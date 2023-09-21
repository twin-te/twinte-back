package apiinterceptor

import (
	"context"
	"errors"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/twin-te/twinte-back/appctx"
	"github.com/twin-te/twinte-back/apperr"
	"github.com/twin-te/twinte-back/idtype"
	authmodule "github.com/twin-te/twinte-back/module/auth"
)

func NewAuthInterceptor(authUseCase authmodule.UseCase) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			cookie, err := (&http.Request{Header: req.Header()}).Cookie("twinte_session")
			if err != nil {
				return next(ctx, req)
			}

			sessionID, err := idtype.NewSessionIDFromString(cookie.Value)
			if err != nil {
				return next(ctx, req)
			}

			user, err := authUseCase.Authenticate(ctx, sessionID)
			if err != nil {
				if errors.Is(err, apperr.ErrUnauthenticated) {
					return next(ctx, req)
				}
				return nil, err
			}

			return next(appctx.SetUser(ctx, user), req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
