package timetableentity

import (
	"github.com/twin-te/twinte-back/idtype"
)

type RegisteredCourse struct {
	ID          idtype.RegisteredCourseID
	UserID      idtype.UserID
	Year        int
	CourseID    *idtype.CourseID
	Name        *string
	Instructors *string
	Cregit      *float64
	Methods     *[]CourseMethod
	Schedules   *[]Schedule
	Memo        string
	Attendance  int
	Absence     int
	Late        int
	TagIDs      idtype.TagIDs
}

func NewRegisteredCourseFromCourse(userID idtype.UserID, course *Course) *RegisteredCourse {
	return &RegisteredCourse{
		ID:       idtype.NewRegisteredCourseID(),
		UserID:   userID,
		CourseID: &course.ID,
		Year:     course.Year,
	}
}
