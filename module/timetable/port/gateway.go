package timetableport

import (
	"context"

	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
)

type Gateway interface {
	// GetCourseWithoutIDsFromKDB returns the latest course data retrieved from kdb.
	GetCourseWithoutIDsFromKDB(ctx context.Context, year shareddomain.AcademicYear) ([]CourseWithoutID, error)
}
