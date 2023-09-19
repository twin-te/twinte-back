package apioauth2

import (
	"errors"
	"net/http"
	"strings"

	"github.com/twin-te/twinte-back/apperr"
	authmodule "github.com/twin-te/twinte-back/module/auth"
	authentity "github.com/twin-te/twinte-back/module/auth/entity"
)

type Handler struct {
	authUseCase authmodule.UseCase
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	var provider authentity.Provider
	switch strings.TrimPrefix(r.URL.Path, "/oauth2/") {
	case "google":
		provider = authentity.ProviderGoogle
	default:
		http.Error(w, "invalid provider", http.StatusInternalServerError)
		return
	}

	state, err := h.authUseCase.GenerateOAuth2State(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url, err := h.authUseCase.GetOAuth2ConsentPageURL(r.Context(), provider, state)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "twinte_oauth2_state",
		Value:    state.String(),
		Path:     "/",
		MaxAge:   180,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, url.String(), http.StatusFound)
}

func (h *Handler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	var provider authentity.Provider
	switch strings.TrimPrefix(r.URL.Path, "/oauth2/callback/") {
	case "google":
		provider = authentity.ProviderGoogle
	default:
		http.Error(w, "invalid provider", http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("twinte_oauth2_state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookieState := cookie.Value
	queryState := r.URL.Query().Get("state")

	if cookieState == "" || queryState == "" || cookieState != queryState {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	code := authentity.NewOAuth2CodeFromString(r.URL.Query().Get("code"))

	socialID, err := h.authUseCase.GetSocialIDFromCode(r.Context(), provider, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := h.authUseCase.SignUpOrLogin(r.Context(), authentity.UserAuthentication{
		Provider: provider,
		SocialID: socialID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "twinte_oauth2_state",
		Path:   "/",
		MaxAge: -1,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "twinte_session",
		Value:    session.ID.String(),
		Path:     "/",
		Expires:  session.ExpiredAt,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "http://localhost:3000", http.StatusFound)
}

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	user, err := h.authUseCase.AuthorizeAuthenticatedUser(r.Context())
	if err != nil {
		if errors.Is(err, apperr.ErrUnauthenticated) {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := h.authUseCase.Logout(r.Context(), user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "twinte_session",
		Path:   "/",
		MaxAge: -1,
	})
	return
}

func NewHandler(authUseCase authmodule.UseCase) *Handler {
	return &Handler{
		authUseCase: authUseCase,
	}
}