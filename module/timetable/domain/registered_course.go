package timetabledomain

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

var (
	ParseAttendance = shareddomain.NewNonNegativeIntParser("attendance")
	ParseAbsence    = shareddomain.NewNonNegativeIntParser("absence")
	ParseLate       = shareddomain.NewNonNegativeIntParser("late")
)

// RegisteredCourse is identified by one of the following fields.
//   - ID
//   - UserID and CourseID ( if CourseID is not nil )
//
// If CourseID is nil, Name, Instructors, Credit, Methods, Schedules are required.
type RegisteredCourse struct {
	ID          idtype.RegisteredCourseID
	UserID      idtype.UserID
	Year        shareddomain.AcademicYear
	CourseID    *idtype.CourseID
	Name        *shareddomain.RequiredString
	Instructors *string
	Credit      *Credit
	Methods     *[]CourseMethod
	Schedules   *[]Schedule
	Memo        string
	Attendance  shareddomain.NonNegativeInt
	Absence     shareddomain.NonNegativeInt
	Late        shareddomain.NonNegativeInt
	TagIDs      []idtype.TagID

	EntityBeforeUpdated *RegisteredCourse
}

func (rc *RegisteredCourse) Clone() *RegisteredCourse {
	ret := lo.ToPtr(*rc)

	if rc.CourseID != nil {
		*ret.CourseID = *rc.CourseID
	}

	if rc.Name != nil {
		*ret.Name = *rc.Name
	}

	if rc.Instructors != nil {
		*ret.Instructors = *rc.Instructors
	}

	if rc.Credit != nil {
		*ret.Credit = *rc.Credit
	}

	if rc.Methods != nil {
		*ret.Methods = base.CopySlice(*rc.Methods)
	}

	if rc.Schedules != nil {
		*ret.Schedules = base.CopySlice(*rc.Schedules)
	}

	ret.TagIDs = base.CopySlice(rc.TagIDs)

	return ret
}

func (rc *RegisteredCourse) BeforeUpdateHook() {
	rc.EntityBeforeUpdated = rc.Clone()
}

type RegisteredCourseDataToUpdate struct {
	Name        *shareddomain.RequiredString
	Instructors *string
	Credit      *Credit
	Methods     *[]CourseMethod
	Schedules   *[]Schedule
	Memo        *string
	Attendance  *shareddomain.NonNegativeInt
	Absence     *shareddomain.NonNegativeInt
	Late        *shareddomain.NonNegativeInt
	TagIDs      *[]idtype.TagID
}

func (rc *RegisteredCourse) Update(data RegisteredCourseDataToUpdate) error {
	if data.Name != nil {
		rc.Name = data.Name
	}

	if data.Instructors != nil {
		rc.Instructors = data.Instructors
	}

	if data.Credit != nil {
		rc.Credit = data.Credit
	}

	if data.Methods != nil {
		rc.Methods = data.Methods
	}

	if data.Schedules != nil {
		rc.Schedules = data.Schedules
	}

	if data.Memo != nil {
		rc.Memo = *data.Memo
	}

	if data.Attendance != nil {
		rc.Attendance = *data.Attendance
	}

	if data.Absence != nil {
		rc.Absence = *data.Absence
	}

	if data.Late != nil {
		rc.Late = *data.Late
	}

	if data.TagIDs != nil {
		rc.TagIDs = *data.TagIDs
	}

	return nil
}

func ConstructRegisteredCourse(fn func(rc *RegisteredCourse) (err error)) (*RegisteredCourse, error) {
	rc := new(RegisteredCourse)
	if err := fn(rc); err != nil {
		return nil, err
	}

	if rc.CourseID == nil && (rc.Name == nil || rc.Instructors == nil || rc.Credit == nil || rc.Methods == nil || rc.Schedules == nil) {
		return nil, fmt.Errorf("the registered course, which does not have course id, must have name, instructors, credit, methods, and schedules. %+v", rc)
	}

	if rc.ID.IsZero() || rc.UserID.IsZero() || rc.Year.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", rc)
	}

	return rc, nil
}
