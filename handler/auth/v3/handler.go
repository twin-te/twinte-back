package authv3

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/twin-te/twinte-back/appenv"
	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
)

func (h *impl) handleOAuth2(c echo.Context) error {
	state := uuid.NewString()
	var url string

	switch c.Param("provider") {
	case "google":
		url = googleOAuth2Config.AuthCodeURL(state)
	case "apple":
		return echo.NewHTTPError(http.StatusNotImplemented, "provider apple is not available")
	case "twitter":
		url = twitterOAuth2Config.AuthCodeURL(state, s256ChallengeOption)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "invalid provider")
	}

	setOAuth2StateCookie(c, state)

	return c.Redirect(http.StatusFound, url)
}

func (h *impl) handleOAuth2Callback(c echo.Context) (err error) {
	defer func() {
		clearOAuth2StateCookie(c)

		if err != nil {
			return
		}

		c.Redirect(http.StatusFound, appenv.AUTH_REDIRECT_URL)
	}()

	cookie, err := c.Cookie(appenv.COOKIE_OAUTH2_STATE_NAME)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "state is not found in cookie")
	}

	cookieState := cookie.Value
	queryState := c.QueryParam("state")

	if cookieState == "" || queryState == "" || cookieState != queryState {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid state")
	}

	code := c.QueryParam("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "authorization code is required")
	}

	var (
		provider authdomain.Provider
		socialID authdomain.SocialID
	)

	switch c.Param("provider") {
	case "google":
		provider = authdomain.ProviderGoogle
		socialID, err = getGoogleSocialID(c.Request().Context(), code)
	case "apple":
		provider = authdomain.ProviderApple
		return echo.NewHTTPError(http.StatusNotImplemented, "provider apple is not available")
	case "twitter":
		provider = authdomain.ProviderTwitter
		socialID, err = getTwitterSocialID(c.Request().Context(), code)
	}
	if err != nil {
		return
	}

	userAuthentication := authdomain.NewUserAuthentication(provider, socialID)

	if _, err := h.accessController.Authenticate(c.Request().Context()); err == nil {
		err = h.authUseCase.AddUserAuthentication(c.Request().Context(), userAuthentication)
		if err != nil {
			return err
		}
	}

	session, err := h.authUseCase.SignUpOrLogin(c.Request().Context(), userAuthentication)
	if err != nil {
		return
	}

	setSessionCookie(c, session)

	return nil
}

func (h *impl) handleIDTokenGoogle(c echo.Context) error {
	idToken := c.QueryParam("token")
	socialID, err := verifyGoogleIDToken(c.Request().Context(), idToken)
	if err != nil {
		return err
	}

	userAuthentication := authdomain.NewUserAuthentication(authdomain.ProviderGoogle, socialID)

	if _, err := h.accessController.Authenticate(c.Request().Context()); err != nil {
		err = h.authUseCase.AddUserAuthentication(c.Request().Context(), userAuthentication)
		if err != nil {
			return err
		}
	} else {
		session, err := h.authUseCase.SignUpOrLogin(c.Request().Context(), userAuthentication)
		if err != nil {
			return err
		}

		setSessionCookie(c, session)
	}

	return c.Redirect(http.StatusFound, appenv.AUTH_REDIRECT_URL)
}

func (h *impl) handleLogout(c echo.Context) error {
	if err := h.authUseCase.Logout(c.Request().Context()); err != nil {
		return err
	}

	clearSessionCookie(c)

	return c.Redirect(http.StatusFound, appenv.AUTH_REDIRECT_URL)
}
