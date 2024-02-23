package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	dbhelper "github.com/twin-te/twinte-back/db/helper"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
	timetableport "github.com/twin-te/twinte-back/module/timetable/port"
	timetablerepository "github.com/twin-te/twinte-back/module/timetable/repository"
)

// checkCoursesCmd represents the check-courses command
var checkCoursesCmd = &cobra.Command{
	Use:   "check-courses",
	Short: "Check if all courses are correctly reconstructed from database",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := dbhelper.NewDB()
		if err != nil {
			log.Fatalln(err)
		}

		timetableRepository := timetablerepository.New(db)

		courses, err := timetableRepository.ListCourses(context.Background(), timetableport.ListCoursesConds{}, sharedport.LockNone)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%d courses are correctly reconstructed from database", len(courses))
	},
}

func init() {
	rootCmd.AddCommand(checkCoursesCmd)
}
