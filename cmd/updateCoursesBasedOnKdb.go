package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/twin-te/twinte-back/appctx"
	dbhelper "github.com/twin-te/twinte-back/db/helper"
	"github.com/twin-te/twinte-back/module/auth/accesscontroller"
	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
	authport "github.com/twin-te/twinte-back/module/auth/port"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	timetablefactory "github.com/twin-te/twinte-back/module/timetable/factory"
	timetablegateway "github.com/twin-te/twinte-back/module/timetable/gateway"
	timetablerepository "github.com/twin-te/twinte-back/module/timetable/repository"
	timetableusecase "github.com/twin-te/twinte-back/module/timetable/usecase"
)

var (
	year            int
	kdbJSONFilePath string
)

// updateCoursesBasedOnKDBCmd represents the update-courses-based-on-kdb command
var updateCoursesBasedOnKDBCmd = &cobra.Command{
	Use:   "update-courses-based-on-kdb",
	Short: "Update courses based on kdb",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := dbhelper.NewDB()
		if err != nil {
			log.Fatalln(err)
		}

		var authRepository authport.Repository
		accessController := accesscontroller.New(authRepository)

		timetableFactory := timetablefactory.New(db)
		timetableGateWay := timetablegateway.New(kdbJSONFilePath)
		timetableRepository := timetablerepository.New(db)
		timetableUseCase := timetableusecase.New(accessController, timetableFactory, timetableGateWay, timetableRepository)

		ctx := appctx.SetActor(context.Background(), authdomain.NewUnknown(authdomain.PermissionExecuteBatchJob))

		year, err := shareddomain.ParseAcademicYear(year)
		if err != nil {
			log.Fatalln(err)
		}

		if err := timetableUseCase.UpdateCoursesBasedOnKDB(ctx, year); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCoursesBasedOnKDBCmd)

	updateCoursesBasedOnKDBCmd.Flags().IntVar(&year, "year", 0, "academic year of courses you want to update")
	updateCoursesBasedOnKDBCmd.Flags().StringVar(&kdbJSONFilePath, "kdb-json-file-path", "", "kdb json file path that is used in timetable gateway")
}
