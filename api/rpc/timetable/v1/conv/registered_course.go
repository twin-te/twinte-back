package timetablev1conv

import (
	"github.com/samber/lo"
	sharedconv "github.com/twin-te/twinte-back/api/rpc/shared/conv"
	timetablev1 "github.com/twin-te/twinte-back/api/rpcgen/timetable/v1"
	"github.com/twin-te/twinte-back/base"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

func ToPBRegisteredCourse(registeredCourse *timetabledomain.RegisteredCourse) (pbRegisteredCourse *timetablev1.RegisteredCourse, err error) {
	pbRegisteredCourse = &timetablev1.RegisteredCourse{
		Id:          sharedconv.ToPBUUID(registeredCourse.ID),
		UserId:      sharedconv.ToPBUUID(registeredCourse.UserID),
		Year:        sharedconv.ToPBAcademicYear(registeredCourse.Year),
		Instructors: registeredCourse.Instructors,
		Memo:        registeredCourse.Memo,
		Attendance:  int32(registeredCourse.Attendance),
		Absence:     int32(registeredCourse.Absence),
		Late:        int32(registeredCourse.Late),
		TagIds:      base.Map(registeredCourse.TagIDs, sharedconv.ToPBUUID[idtype.TagID]),
	}

	if registeredCourse.CourseID != nil {
		pbRegisteredCourse.CourseId = sharedconv.ToPBUUID(*registeredCourse.CourseID)
	}

	if registeredCourse.Name != nil {
		pbRegisteredCourse.Name = lo.ToPtr(registeredCourse.Name.String())
	}

	if registeredCourse.Credit != nil {
		pbRegisteredCourse.Credit = lo.ToPtr(registeredCourse.Credit.String())
	}

	if registeredCourse.Methods != nil {
		courseMethodList := new(timetablev1.CourseMethodList)
		courseMethodList.Values, err = base.MapWithErr(*registeredCourse.Methods, ToPBCourseMethod)
		if err != nil {
			return
		}
		pbRegisteredCourse.Methods = courseMethodList
	}

	if registeredCourse.Schedules != nil {
		scheduleList := new(timetablev1.ScheduleList)
		scheduleList.Values, err = base.MapWithErr(*registeredCourse.Schedules, ToPBSchedule)
		if err != nil {
			return
		}
		pbRegisteredCourse.Schedules = scheduleList
	}

	return
}
