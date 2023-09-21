package cmd

import (
	"database/sql"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/twin-te/twinte-back/api/gen/apigenconnect"
	apiinterceptor "github.com/twin-te/twinte-back/api/interceptor"
	apioauth2 "github.com/twin-te/twinte-back/api/oauth2"
	apiservice "github.com/twin-te/twinte-back/api/service"
	"github.com/twin-te/twinte-back/appenv"
	authgateway "github.com/twin-te/twinte-back/module/auth/gateway"
	authrepository "github.com/twin-te/twinte-back/module/auth/repository"
	authusecase "github.com/twin-te/twinte-back/module/auth/usecase"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve the api server",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("postgres", appenv.DB_URL)
		if err != nil {
			panic(err)
		}

		authGateway := authgateway.New()
		authRepository := authrepository.New(db)
		authUseCase := authusecase.New(authGateway, authRepository)
		authAPIService := apiservice.NewAuthService(authUseCase)

		apiOAuth2Handler := apioauth2.NewHandler(authUseCase)

		intercepters := connect.WithInterceptors(apiinterceptor.NewAuthInterceptor(authUseCase))

		mux := http.NewServeMux()

		mux.HandleFunc("/oauth2/logout", apiOAuth2Handler.HandleLogout)
		mux.HandleFunc("/oauth2/callback/", apiOAuth2Handler.HandleCallback)
		mux.HandleFunc("/oauth2/", apiOAuth2Handler.Handle)

		mux.Handle(apigenconnect.NewAuthServiceHandler(authAPIService, intercepters))

		http.ListenAndServe(
			":8080",
			// Use h2c so we can serve HTTP/2 without TLS.
			h2c.NewHandler(mux, &http2.Server{}),
		)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
