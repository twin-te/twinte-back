package restv4

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/appctx"
	"github.com/twin-te/twinte-back/appenv"
	"github.com/twin-te/twinte-back/apperr"
	authmodule "github.com/twin-te/twinte-back/module/auth"
	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
	autherr "github.com/twin-te/twinte-back/module/auth/err"
	sharederr "github.com/twin-te/twinte-back/module/shared/err"
	"golang.org/x/oauth2"
)

var twitterOAuth2Config = &oauth2.Config{
	ClientID:     appenv.OAUTH2_TWITTER_CLIENT_ID,
	ClientSecret: appenv.OAUTH2_TWITTER_CLIENT_SECRET,
	Endpoint: oauth2.Endpoint{
		AuthURL:   "https://twitter.com/i/oauth2/authorize",
		TokenURL:  "https://api.twitter.com/2/oauth2/token",
		AuthStyle: oauth2.AuthStyleInHeader,
	},
	RedirectURL: appenv.OAUTH2_TWITTER_CALLBACK_URL,
	Scopes:      []string{"users.read", "tweet.read"},
}

var (
	verifier            = oauth2.GenerateVerifier()
	verifierOption      = oauth2.VerifierOption(verifier)
	s256ChallengeOption = oauth2.S256ChallengeOption(verifier)
)

type impl struct {
	authUseCase authmodule.UseCase
}

func (h *impl) handleOAuth2(w http.ResponseWriter, r *http.Request) {
	state := uuid.NewString()
	var url string

	switch strings.TrimPrefix(r.URL.Path, "/oauth2/") {
	case "google":
		http.Error(w, "provider google is not available", http.StatusNotImplemented)
	case "apple":
		http.Error(w, "provider apple is not available", http.StatusNotImplemented)
		return
	case "twitter":
		url = twitterOAuth2Config.AuthCodeURL(state, s256ChallengeOption)
		http.Error(w, "provider twitter is not available", http.StatusNotImplemented)
	default:
		http.Error(w, "invalid provider", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     appenv.COOKIE_OAUTH2_STATE_NAME,
		Value:    state,
		Path:     "/",
		MaxAge:   appenv.COOKIE_OAUTH2_STATE_MAX_AGE,
		Secure:   appenv.COOKIE_SECURE,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *impl) handleOAuth2Callback(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(appenv.COOKIE_OAUTH2_STATE_NAME)
	if err != nil {
		http.Error(w, "oauth2 state is not found in cookie", http.StatusBadRequest)
		return
	}

	cookieState := cookie.Value
	queryState := r.URL.Query().Get("state")

	if cookieState == "" || queryState == "" || cookieState != queryState {
		http.Error(w, "invalid oauth2 state", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "code is required", http.StatusBadRequest)
		return
	}

	var (
		provider authdomain.Provider
		socialID authdomain.SocialID
	)

	switch strings.TrimPrefix(r.URL.Path, "/oauth2/callback/") {
	case "google":
		provider = authdomain.ProviderGoogle
		socialID, err = h.getGoogleSocialID(r.Context(), code)
	case "apple":
		provider = authdomain.ProviderApple
		http.Error(w, "provider apple is not available", http.StatusNotImplemented)
		return
	case "twitter":
		provider = authdomain.ProviderTwitter
		socialID, err = h.getTwitterSocialID(r.Context(), code)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userAuthentication := authdomain.NewUserAuthentication(provider, socialID)

	if actor, ok := appctx.GetActor(r.Context()); ok && actor.AuthNUser() != nil {
		err = h.authUseCase.AddUserAuthentication(r.Context(), userAuthentication)
		if err != nil {
			if aerr, ok := lo.ErrorsAs[*apperr.Error](err); ok {
				if aerr.Code == sharederr.CodeAlreadyExists {
					http.Error(w, aerr.Error(), http.StatusBadRequest)
					return
				}

				if aerr.Code == autherr.CodeMultipleAuthenticationOfSameProvider {
					http.Error(w, aerr.Error(), http.StatusBadRequest)
					return
				}
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		session, err := h.authUseCase.SignUpOrLogin(r.Context(), userAuthentication)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:   appenv.COOKIE_OAUTH2_STATE_NAME,
			Path:   "/",
			MaxAge: -1,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     appenv.COOKIE_SESSION_NAME,
			Value:    session.ID.String(),
			Path:     "/",
			Expires:  session.ExpiredAt,
			Secure:   appenv.COOKIE_SECURE,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
	}

	http.Redirect(w, r, appenv.AUTH_REDIRECT_URL, http.StatusPermanentRedirect)
}

func (h *impl) getGoogleSocialID(ctx context.Context, code string) (socialID authdomain.SocialID, err error) {
	return
}

func (h *impl) getTwitterSocialID(ctx context.Context, code string) (socialID authdomain.SocialID, err error) {
	token, err := twitterOAuth2Config.Exchange(ctx, code, verifierOption)
	if err != nil {
		return
	}

	client := twitterOAuth2Config.Client(ctx, token)
	resp, err := client.Get("https://api.twitter.com/2/users/me")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	v := &struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}{}

	if err = json.Unmarshal(body, v); err != nil {
		return
	}

	return authdomain.ParseSocialID(v.Data.ID)
}

func (h *impl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/v4/oauth2/callback/":
		h.handleOAuth2Callback(w, r)
	case "/v4/oauth2/":
		h.handleOAuth2(w, r)
	default:
		http.NotFound(w, r)
	}
}

func New(authUseCase authmodule.UseCase) (string, *impl) {
	return "/v4/", &impl{
		authUseCase: authUseCase,
	}
}
