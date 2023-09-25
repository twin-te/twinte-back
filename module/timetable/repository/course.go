package timetablerepository

import (
	"context"

	dbmodel "github.com/twin-te/twinte-back/db/models"
	timetableentity "github.com/twin-te/twinte-back/module/timetable/entity"
)

func (r *Impl) SaveCourses(ctx context.Context, courses []*timetableentity.Course) error {
	c := &dbmodel.Course{}

	s := &dbmodel.CourseSchedule{}
	g := &dbmodel.CourseRecommendedGrade{}

	return nil
}
