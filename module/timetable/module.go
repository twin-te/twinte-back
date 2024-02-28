package timetablemodule

import (
	"context"

	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

// UseCase represents application specific business rules.
//
// The error codes for authentication and authorization failures are not stated explicitly.
type UseCase interface {
	// GetCoursesByIDs returns the courses specified by the given ids.
	// If any of the target courses is not found, no error occurs.
	//
	// [Authentication] not required
	GetCoursesByIDs(ctx context.Context, ids []idtype.CourseID) ([]*timetabledomain.Course, error)

	// GetCoursesByCodes returns the courses specified by the given year and codes.
	// If any of the target courses is not found, no error occurs.
	//
	// [Authentication] not required
	GetCoursesByCodes(ctx context.Context, year shareddomain.AcademicYear, codes []timetabledomain.Code) ([]*timetabledomain.Course, error)

	// SearchCourses returns the courses satisfied with the conditions.
	//
	// [Authentication] not required
	SearchCourses(ctx context.Context, in SearchCoursesIn) ([]*timetabledomain.Course, error)

	// CreateRegisteredCoursesByCodes creates new registered courses by the given year and codes.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - shared.AlreadyExists
	//   - shared.NotFound
	CreateRegisteredCoursesByCodes(ctx context.Context, year shareddomain.AcademicYear, codes []timetabledomain.Code) ([]*timetabledomain.RegisteredCourse, error)

	// CreateRegisteredCourseManually creates a new registered course mannually.
	//
	// [Authentication] required
	CreateRegisteredCourseManually(ctx context.Context, in CreateRegisteredCourseManuallyIn) (*timetabledomain.RegisteredCourse, error)

	// GetRegisteredCourseByID returns the registered course specified by the given id.
	//
	// [Authentication] required
	GetRegisteredCourseByID(ctx context.Context, id idtype.RegisteredCourseID) (*timetabledomain.RegisteredCourse, error)

	// GetRegisteredCourses returns the registered courses.
	//
	// [Authentication] required
	GetRegisteredCourses(ctx context.Context, year *shareddomain.AcademicYear) ([]*timetabledomain.RegisteredCourse, error)

	// UpdateRegisteredCourse updates registered course specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - shared.NotFound
	//   - shared.InvalidArgument ( if the user does not have tags specified by the given TagIDs )
	UpdateRegisteredCourse(ctx context.Context, in UpdateRegisteredCourseIn) (*timetabledomain.RegisteredCourse, error)

	// DeleteRegisteredCourse deletes registered course specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - shared.NotFound
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
	//   - shared.NotFound
	UpdateTag(ctx context.Context, in UpdateTagIn) (*timetabledomain.Tag, error)

	// DeleteTag deletes the tag specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - shared.NotFound
	DeleteTag(ctx context.Context, id idtype.TagID) error

	// RearrangeTags rearranges the tags.
	// Please specify all tags associated with the user.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - shared.InvalidArgument
	RearrangeTags(ctx context.Context, tagIDs []idtype.TagID) error

	// UpdateCoursesBasedOnKDB retrieves data about courses from kdb and updates courses.
	//
	// [Authentication] not required
	//
	// [Permission]
	//   - PermissionExecuteBatchJob
	//
	// [Error Code]
	//   - shared.NotFound
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
