package timetablemodule

import (
	"context"

	"github.com/twin-te/twinte-back/idtype"
	timetableentity "github.com/twin-te/twinte-back/module/timetable/entity"
)

type UseCase interface {
	// UpdateCoursesBasedOnKDB(ctx context.Context) error

	// authenticated user

	// SearchCourses(ctx context.Context, in SearchCoursesIn) error

	CreateRegisteredCoursesByCodes(ctx context.Context, year int, codes []timetableentity.Code) error
	CreateRegisteredCourseManually(ctx context.Context, in CreateCourseManuallyIn) error
	GetRegisteredCourses(ctx context.Context, year *int) ([]*timetableentity.RegisteredCourse, error)
	DeleteRegisteredCourse(ctx context.Context, id idtype.RegisteredCourseID) error

	CreateTag(ctx context.Context, name string) error
	GetTags(ctx context.Context) ([]*timetableentity.Tag, error)
	UpdateTagName(ctx context.Context, id idtype.TagID, name string) error
	DeleteTag(ctx context.Context, id idtype.TagID) error
	RearrangeTags(ctx context.Context, tagIDs idtype.TagIDs) error
}

// type SearchCoursesIn struct {
// 	Keyword   string
// 	Code      string
// 	IsBlanck  bool
// 	Schedules []timetableentity.Schedule
// }

type CreateCourseManuallyIn struct {
	Year        int
	Name        string
	Instructors string
	Cregit      float64
	Methods     []timetableentity.CourseMethod
	Schedules   []timetableentity.Schedule
}
