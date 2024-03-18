package middleware

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/appenv"
	"github.com/twin-te/twinte-back/apperr"
	authmodule "github.com/twin-te/twinte-back/module/auth"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharederr "github.com/twin-te/twinte-back/module/shared/err"
)

func NewEchoErrorHandler() func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			err = next(c)

			if err == nil {
				return
			}

			if aerr, ok := apperr.As(err); ok {
				switch aerr.Code {
				case sharederr.CodeInvalidArgument:
					return echo.NewHTTPError(http.StatusBadRequest, aerr.Message)
				case sharederr.CodeUnauthenticated:
					return echo.ErrUnauthorized
				case sharederr.CodeUnauthorized:
					return echo.ErrForbidden
				}
			}

			if httpError, ok := err.(*echo.HTTPError); !ok || httpError.Code >= 500 {
				log.Printf("unexpected error occurred: %+v", err)
			}

			return
		}
	}
}

func NewEchoWithActor(accessController authmodule.AccessController) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sessionID, ok := getSessionIDFromEchoContext(c)
			ctx, err := accessController.WithActor(c.Request().Context(), lo.Ternary(ok, &sessionID, nil))
			if err != nil {
				return err
			}

			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

func getSessionIDFromEchoContext(c echo.Context) (sessionID idtype.SessionID, ok bool) {
	cookie, err := c.Cookie(appenv.COOKIE_SESSION_NAME)
	if err != nil {
		return
	}

	sessionID, err = idtype.ParseSessionID(cookie.Value)
	if err != nil {
		return
	}

	return sessionID, true
}
