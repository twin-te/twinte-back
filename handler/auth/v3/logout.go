package authv3

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/twin-te/twinte-back/appenv"
)

func (h *impl) handleLogout(c echo.Context) error {
	if err := h.authUseCase.Logout(c.Request().Context()); err != nil {
		return err
	}

	clearSessionCookie(c)

	return c.Redirect(http.StatusFound, appenv.AUTH_REDIRECT_URL)
}
