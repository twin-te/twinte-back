package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/twin-te/twinte-back/appenv"
	dbhelper "github.com/twin-te/twinte-back/db/helper"
	"github.com/twin-te/twinte-back/handler"
	announcementfactory "github.com/twin-te/twinte-back/module/announcement/factory"
	announcementrepository "github.com/twin-te/twinte-back/module/announcement/repository"
	announcementusecase "github.com/twin-te/twinte-back/module/announcement/usecase"
	"github.com/twin-te/twinte-back/module/auth/accesscontroller"
	authfactory "github.com/twin-te/twinte-back/module/auth/factory"
	authrepository "github.com/twin-te/twinte-back/module/auth/repository"
	authusecase "github.com/twin-te/twinte-back/module/auth/usecase"
	donationmodule "github.com/twin-te/twinte-back/module/donation"
	schoolcalendarrepository "github.com/twin-te/twinte-back/module/schoolcalendar/repository"
	schoolcalendarusecase "github.com/twin-te/twinte-back/module/schoolcalendar/usecase"
	timetablefactory "github.com/twin-te/twinte-back/module/timetable/factory"
	timetablegateway "github.com/twin-te/twinte-back/module/timetable/gateway"
	timetablerepository "github.com/twin-te/twinte-back/module/timetable/repository"
	timetableusecase "github.com/twin-te/twinte-back/module/timetable/usecase"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := dbhelper.NewDB()
		if err != nil {
			log.Fatalln(err)
		}

		nowFunc := func() time.Time { return time.Now().Truncate(time.Microsecond) }

		authFactory := authfactory.New(nowFunc)
		authRepository := authrepository.New(db)
		accessController := accesscontroller.New(authRepository)
		authUseCase := authusecase.New(accessController, authFactory, authRepository)

		announcementFactory := announcementfactory.New(nowFunc)
		announcementRepository := announcementrepository.New(db)
		announcementUsecase := announcementusecase.New(accessController, announcementFactory, announcementRepository)

		var dummyDonationUseCase donationmodule.UseCase
		// donationFactory := donationfactory.New()
		// donationGateway := donationgateway.New()
		// donationRepository := donationrepository.New(db)
		// donationUseCase := donationusecase.New(accessController, donationFactory, donationGateway, donationRepository)

		schoolcalendarRepository := schoolcalendarrepository.New()
		schoolcalendarUseCase := schoolcalendarusecase.New(accessController, schoolcalendarRepository)

		timetableFactory := timetablefactory.New(db)
		timetableGateWay := timetablegateway.New("")
		timetableRepository := timetablerepository.New(db)
		timetableUseCase := timetableusecase.New(accessController, timetableFactory, timetableGateWay, timetableRepository)

		h := handler.New(
			accessController,
			announcementUsecase,
			authUseCase,
			dummyDonationUseCase,
			schoolcalendarUseCase,
			timetableUseCase,
		)

		mux := http.NewServeMux()
		mux.Handle("/", h)

		if err := http.ListenAndServe(appenv.ADDR, mux); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
