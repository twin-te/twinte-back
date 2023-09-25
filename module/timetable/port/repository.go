package timetableport

import (
	"context"

	"github.com/twin-te/twinte-back/idtype"
	timetableentity "github.com/twin-te/twinte-back/module/timetable/entity"
)

type Repository interface {
	Transaction(ctx context.Context, fc func(rtx Repository) error, readOnly bool) error

	SaveCourses(ctx context.Context, courses []*timetableentity.Course) error
	ListCourses(ctx context.Context, conds ListCoursesConds) ([]*timetableentity.Course, error)

	SaveRegisteredCourse(ctx context.Context, registeredCourse *timetableentity.RegisteredCourse) error
	SaveRegisteredCourses(ctx context.Context, registeredCourses []*timetableentity.RegisteredCourse) error
	FindRegisteredCourse(ctx context.Context, conds FindRegisteredCourseConds) (*timetableentity.RegisteredCourse, error)
	ListRegisteredCourses(ctx context.Context, conds ListRegisteredCoursesConds) ([]*timetableentity.RegisteredCourse, error)
	DeleteRegisteredCourse(ctx context.Context, id idtype.RegisteredCourseID) error

	SaveTag(ctx context.Context, tag *timetableentity.Tag) error
	SaveTags(ctx context.Context, tags []*timetableentity.Tag) error
	FindTag(ctx context.Context, conds FindTagConds) (*timetableentity.Tag, error)
	ListTags(ctx context.Context, conds ListTagsConds) ([]*timetableentity.Tag, error)
	DeleteTag(ctx context.Context, id idtype.TagID) error
}

type ListCoursesConds struct {
	IDs   *idtype.CourseIDs
	Year  *int
	Codes *[]timetableentity.Code
}

type FindRegisteredCourseConds struct {
	ID     *idtype.RegisteredCourseID
	UserID *idtype.UserID
}

type ListRegisteredCoursesConds struct {
	UserID *idtype.UserID
	Year   *int
}

type FindTagConds struct {
	ID     *idtype.TagID
	UserID *idtype.UserID
}

type ListTagsConds struct {
	UserID *idtype.UserID
}
