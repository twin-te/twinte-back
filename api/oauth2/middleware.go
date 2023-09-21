package apioauth2

import (
	"errors"
	"net/http"

	"github.com/twin-te/twinte-back/appctx"
	"github.com/twin-te/twinte-back/appenv"
	"github.com/twin-te/twinte-back/apperr"
	"github.com/twin-te/twinte-back/idtype"
	authmodule "github.com/twin-te/twinte-back/module/auth"
)

func NewAuthMiddleware(authUseCase authmodule.UseCase) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(appenv.COOKIE_SESSION_NAME)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			sessionID, err := idtype.NewSessionIDFromString(cookie.Value)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			user, err := authUseCase.Authenticate(r.Context(), sessionID)
			if err != nil {
				if errors.Is(err, apperr.ErrUnauthenticated) {
					next.ServeHTTP(w, r)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			next.ServeHTTP(w, r.WithContext(appctx.SetUser(r.Context(), user)))
		})
	}
}
