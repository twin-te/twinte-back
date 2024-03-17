package timetablemodule

import (
	"context"

	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

// UseCase represents application specific business rules.
//
// The following error codes are not stated explicitly in the each method, but may be returned.
//   - shared.InvalidArgument
//   - shared.Unauthenticated
//   - shared.Unauthorized
type UseCase interface {
	// GetCoursesByIDs returns the courses specified by the given ids.
	// If any of the target courses is not found, no error will not be returned.
	//
	// [Authentication] not required
	GetCoursesByIDs(ctx context.Context, ids []idtype.CourseID) ([]*timetabledomain.Course, error)

	// GetCoursesByCodes returns the courses specified by the given year and codes.
	// If any of the target courses is not found, no error will not be returned.
	//
	// [Authentication] not required
	GetCoursesByCodes(ctx context.Context, year shareddomain.AcademicYear, codes []timetabledomain.Code) ([]*timetabledomain.Course, error)

	// SearchCourses returns the courses satisfied with the conditions.
	//
	// [Authentication] not required
	SearchCourses(ctx context.Context, in SearchCoursesIn) ([]*timetabledomain.Course, error)

	// CreateRegisteredCoursesByCodes creates new registered courses by the given year and codes.
	// And it returns the registered courses, each of which has the course association loaded if it has the based course.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - timetable.CourseNotFound
	//   - timetable.RegisteredCourseAlreadyExists
	CreateRegisteredCoursesByCodes(ctx context.Context, year shareddomain.AcademicYear, codes []timetabledomain.Code) ([]*timetabledomain.RegisteredCourse, error)

	// CreateRegisteredCourseManually creates a new registered course mannually.
	// And it returns the registered course, which has the course association loaded if it has the based course.
	//
	// [Authentication] required
	CreateRegisteredCourseManually(ctx context.Context, in CreateRegisteredCourseManuallyIn) (*timetabledomain.RegisteredCourse, error)

	// GetRegisteredCourseByID returns the registered course specified by the given id.
	// The returned registered course has the course association loaded if it has the based course.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - timetable.RegisteredCourseNotFound
	GetRegisteredCourseByID(ctx context.Context, id idtype.RegisteredCourseID) (*timetabledomain.RegisteredCourse, error)

	// GetRegisteredCourses returns the registered courses.
	// Each of the returned registered courses has the course association loaded if it has the based course.
	//
	// [Authentication] required
	GetRegisteredCourses(ctx context.Context, year *shareddomain.AcademicYear) ([]*timetabledomain.RegisteredCourse, error)

	// UpdateRegisteredCourse updates registered course specified by the given id.
	// And it returns the registered course, which has the course association loaded if it has the based course.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - timetable.RegisteredCourseNotFound
	UpdateRegisteredCourse(ctx context.Context, in UpdateRegisteredCourseIn) (*timetabledomain.RegisteredCourse, error)

	// DeleteRegisteredCourse deletes registered course specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - timetable.RegisteredCourseNotFound
	DeleteRegisteredCourse(ctx context.Context, id idtype.RegisteredCourseID) error

	// CreateTag creates a new tag.
	//
	// [Authentication] required
	CreateTag(ctx context.Context, name shareddomain.RequiredString) (tag *timetabledomain.Tag, err error)

	// GetTags returns the tags.
	//
	// [Authentication] required
	GetTags(ctx context.Context) ([]*timetabledomain.Tag, error)

	// UpdateTag updates the tag specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - timetable.TagNotFound
	UpdateTag(ctx context.Context, in UpdateTagIn) (*timetabledomain.Tag, error)

	// DeleteTag deletes the tag specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - timetable.TagNotFound
	DeleteTag(ctx context.Context, id idtype.TagID) error

	// RearrangeTags rearranges the tags.
	// Please specify all tag ids associated with the user.
	//
	// [Authentication] required
	RearrangeTags(ctx context.Context, tagIDs []idtype.TagID) error

	// UpdateCoursesBasedOnKDB retrieves data about courses from kdb and updates courses.
	//
	// [Authentication] not required
	//
	// [Permission]
	//   - PermissionExecuteBatchJob
	UpdateCoursesBasedOnKDB(ctx context.Context, year shareddomain.AcademicYear) error
}

type SearchCoursesIn struct {
	Year         shareddomain.AcademicYear
	Keywords     []string // return the courses whose name contains all specified keywords
	CodePrefixes struct {
		Included []string // return the courses whose code has all specified prefixes.
		Excluded []string // return the courses whose code does not have all specified prefixes.
	}
	Schedules struct {
		FullyIncluded       []timetabledomain.Schedule // return the courses whose schedules are fully included in the specified schedules
		PartiallyOverlapped []timetabledomain.Schedule // return the courses whose schedules are partially overlapped with the specified schedules
	}
	Limit  int
	Offset int
}

type CreateRegisteredCourseManuallyIn struct {
	Year        shareddomain.AcademicYear
	Name        shareddomain.RequiredString
	Instructors string
	Credit      timetabledomain.Credit
	Methods     []timetabledomain.CourseMethod
	Schedules   []timetabledomain.Schedule
}

type UpdateRegisteredCourseIn struct {
	ID          idtype.RegisteredCourseID
	Name        *shareddomain.RequiredString
	Instructors *string
	Credit      *timetabledomain.Credit
	Methods     *[]timetabledomain.CourseMethod
	Schedules   *[]timetabledomain.Schedule
	Memo        *string
	Attendance  *shareddomain.NonNegativeInt
	Absence     *shareddomain.NonNegativeInt
	Late        *shareddomain.NonNegativeInt
	TagIDs      *[]idtype.TagID
}

type UpdateTagIn struct {
	ID   idtype.TagID
	Name *shareddomain.RequiredString
}
