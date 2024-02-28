package timetableport

import (
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

type Factory interface {
	NewCourse(courseWithoutID CourseWithoutID) (*timetabledomain.Course, error)

	NewRegisteredCourseFromCourse(userID idtype.UserID, course *timetabledomain.Course) (*timetabledomain.RegisteredCourse, error)

	NewRegisteredCourseMannualy(
		userID idtype.UserID,
		year shareddomain.AcademicYear,
		name shareddomain.RequiredString,
		instructors string,
		credit timetabledomain.Credit,
		methods []timetabledomain.CourseMethod,
		schedules []timetabledomain.Schedule,
	) (*timetabledomain.RegisteredCourse, error)

	NewTag(
		userID idtype.UserID,
		name shareddomain.RequiredString,
	) (*timetabledomain.Tag, error)
}
