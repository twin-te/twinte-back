package authv3

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/twin-te/twinte-back/appenv"
	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
)

func setOAuth2StateCookie(c echo.Context, state string) {
	c.SetCookie(&http.Cookie{
		Name:     appenv.COOKIE_OAUTH2_STATE_NAME,
		Value:    state,
		Path:     "/",
		MaxAge:   appenv.COOKIE_OAUTH2_STATE_MAX_AGE,
		Secure:   appenv.COOKIE_SECURE,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func clearOAuth2StateCookie(c echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:   appenv.COOKIE_OAUTH2_STATE_NAME,
		Path:   "/",
		MaxAge: -1,
	})
}

func setSessionCookie(c echo.Context, session *authdomain.Session) {
	c.SetCookie(&http.Cookie{
		Name:     appenv.COOKIE_SESSION_NAME,
		Value:    session.ID.String(),
		Path:     "/",
		Expires:  session.ExpiredAt,
		Secure:   appenv.COOKIE_SECURE,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func clearSessionCookie(c echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:   appenv.COOKIE_SESSION_NAME,
		Path:   "/",
		MaxAge: -1,
	})
}
