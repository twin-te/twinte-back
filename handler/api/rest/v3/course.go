package restv3

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/twin-te/twinte-back/apperr"
	"github.com/twin-te/twinte-back/base"
	"github.com/twin-te/twinte-back/handler/api/rest/v3/openapi"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	sharederr "github.com/twin-te/twinte-back/module/shared/err"
	timetablemodule "github.com/twin-te/twinte-back/module/timetable"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

func fromApiCredit(apiCredit float32) (timetabledomain.Credit, error) {
	return timetabledomain.ParseCredit(fmt.Sprintf("%.1f", apiCredit))
}

func toApiCredit(credit timetabledomain.Credit) (float32, error) {
	c, err := strconv.ParseFloat(credit.String(), 32)
	return float32(c), err
}

func fromApiCourseModule(apiModule openapi.CourseModule) (timetabledomain.Module, error) {
	switch apiModule {
	case openapi.CourseModuleSpringA:
		return timetabledomain.ModuleSpringA, nil
	case openapi.CourseModuleSpringB:
		return timetabledomain.ModuleSpringB, nil
	case openapi.CourseModuleSpringC:
		return timetabledomain.ModuleSpringC, nil
	case openapi.CourseModuleSummerVacation:
		return timetabledomain.ModuleSummerVacation, nil
	case openapi.CourseModuleFallA:
		return timetabledomain.ModuleFallA, nil
	case openapi.CourseModuleFallB:
		return timetabledomain.ModuleFallB, nil
	case openapi.CourseModuleFallC:
		return timetabledomain.ModuleFallC, nil
	case openapi.CourseModuleSpringVacation:
		return timetabledomain.ModuleSpringVacation, nil
	}
	return 0, fmt.Errorf("invalid %#v", apiModule)
}

func toApiCourseModule(module timetabledomain.Module) (openapi.CourseModule, error) {
	switch module {
	case timetabledomain.ModuleSpringA:
		return openapi.CourseModuleSpringA, nil
	case timetabledomain.ModuleSpringB:
		return openapi.CourseModuleSpringB, nil
	case timetabledomain.ModuleSpringC:
		return openapi.CourseModuleSpringC, nil
	case timetabledomain.ModuleSummerVacation:
		return openapi.CourseModuleSummerVacation, nil
	case timetabledomain.ModuleFallA:
		return openapi.CourseModuleFallA, nil
	case timetabledomain.ModuleFallB:
		return openapi.CourseModuleFallB, nil
	case timetabledomain.ModuleFallC:
		return openapi.CourseModuleFallC, nil
	case timetabledomain.ModuleSpringVacation:
		return openapi.CourseModuleSpringVacation, nil
	}
	return "", fmt.Errorf("invalid %#v", module)
}

func fromApiDay(apiDay openapi.CourseDay) (timetabledomain.Day, error) {
	switch apiDay {
	case openapi.CourseDaySun:
		return timetabledomain.DaySun, nil
	case openapi.CourseDayMon:
		return timetabledomain.DayMon, nil
	case openapi.CourseDayTue:
		return timetabledomain.DayTue, nil
	case openapi.CourseDayWed:
		return timetabledomain.DayWed, nil
	case openapi.CourseDayThu:
		return timetabledomain.DayThu, nil
	case openapi.CourseDayFri:
		return timetabledomain.DayFri, nil
	case openapi.CourseDaySat:
		return timetabledomain.DaySat, nil
	case openapi.CourseDayIntensive:
		return timetabledomain.DayIntensive, nil
	case openapi.CourseDayAppointment:
		return timetabledomain.DayAppointment, nil
	case openapi.CourseDayAnyTime:
		return timetabledomain.DayAnyTime, nil
	}
	return 0, fmt.Errorf("invalid %#v", apiDay)
}

func toApiDay(day timetabledomain.Day) (openapi.CourseDay, error) {
	switch day {
	case timetabledomain.DaySun:
		return openapi.CourseDaySun, nil
	case timetabledomain.DayMon:
		return openapi.CourseDayMon, nil
	case timetabledomain.DayTue:
		return openapi.CourseDayTue, nil
	case timetabledomain.DayWed:
		return openapi.CourseDayWed, nil
	case timetabledomain.DayThu:
		return openapi.CourseDayThu, nil
	case timetabledomain.DayFri:
		return openapi.CourseDayFri, nil
	case timetabledomain.DaySat:
		return openapi.CourseDaySat, nil
	case timetabledomain.DayIntensive:
		return openapi.CourseDayIntensive, nil
	case timetabledomain.DayAppointment:
		return openapi.CourseDayAppointment, nil
	case timetabledomain.DayAnyTime:
		return openapi.CourseDayAnyTime, nil
	}
	return "", fmt.Errorf("invalid %#v", day)
}

func fromApiCourseMethod(apiCourseMethod openapi.CourseMethod) (timetabledomain.CourseMethod, error) {
	switch apiCourseMethod {
	case openapi.Asynchronous:
		return timetabledomain.CourseMethodOnlineAsynchronous, nil
	case openapi.Synchronous:
		return timetabledomain.CourseMethodOnlineSynchronous, nil
	case openapi.FaceToFace:
		return timetabledomain.CourseMethodFaceToFace, nil
	case openapi.Others:
		return timetabledomain.CourseMethodOthers, nil
	}
	return 0, fmt.Errorf("invalid %#v", apiCourseMethod)
}

func toApiCourseMethod(courseMethod timetabledomain.CourseMethod) (openapi.CourseMethod, error) {
	switch courseMethod {
	case timetabledomain.CourseMethodOnlineAsynchronous:
		return openapi.Asynchronous, nil
	case timetabledomain.CourseMethodOnlineSynchronous:
		return openapi.Synchronous, nil
	case timetabledomain.CourseMethodFaceToFace:
		return openapi.FaceToFace, nil
	case timetabledomain.CourseMethodOthers:
		return openapi.Others, nil
	}
	return "", fmt.Errorf("invalid %#v", courseMethod)
}

func fromApiCourseSchedule(apiSchedule openapi.CourseSchedule) (timetabledomain.Schedule, error) {
	return timetabledomain.ConstructSchedule(func() (schedule timetabledomain.Schedule, err error) {
		schedule.Module, err = fromApiCourseModule(apiSchedule.Module)
		if err != nil {
			return
		}

		schedule.Day, err = fromApiDay(apiSchedule.Day)
		if err != nil {
			return
		}

		if schedule.Day.IsNormal() {
			schedule.Period, err = timetabledomain.ParsePeriod(apiSchedule.Period)
			if err != nil {
				return
			}
		}

		schedule.Rooms = apiSchedule.Room

		return
	})
}

func toApiCourseSchedule(schedule timetabledomain.Schedule) (ret openapi.CourseSchedule, err error) {
	ret.Module, err = toApiCourseModule(schedule.Module)
	if err != nil {
		return
	}

	ret.Day, err = toApiDay(schedule.Day)
	if err != nil {
		return
	}

	ret.Period = schedule.Period.Int()
	ret.Room = schedule.Rooms

	return
}

func toApiCourse(course *timetabledomain.Course) (ret openapi.Course, err error) {
	ret = openapi.Course{
		Code:          course.Code.String(),
		HasParseError: course.HasParseError,
		Id:            toApiUUID(course.ID),
		Instructor:    course.Instructors,
		IsAnnual:      course.IsAnnual,
		Name:          course.Name.String(),
		Overview:      course.Overview,
		Remarks:       course.Remarks,
		Year:          course.Year.Int(),
	}

	ret.Credit, err = toApiCredit(course.Credit)
	if err != nil {
		return
	}

	ret.Methods, err = base.MapWithErr(course.Methods, toApiCourseMethod)
	if err != nil {
		return
	}

	ret.RecommendedGrades = base.Map(course.RecommendedGrades, func(grade timetabledomain.RecommendedGrade) int {
		return grade.Int()
	})

	ret.Schedules, err = base.MapWithErr(course.Schedules, toApiCourseSchedule)
	if err != nil {
		return
	}

	return
}

// 指定した複数の講義を取得する
// (GET /courses)
func (h *impl) GetCourses(ctx context.Context, request openapi.GetCoursesRequestObject) (res openapi.GetCoursesResponseObject, err error) {
	year, err := shareddomain.ParseAcademicYear(request.Params.Year)
	if err != nil {
		return
	}

	codes, err := base.MapWithErr(strings.Split(request.Params.Codes, ","), timetabledomain.ParseCode)
	if err != nil {
		return
	}

	courses, err := h.timetableUseCase.GetCoursesByCodes(ctx, year, codes)
	if err != nil {
		return
	}

	apiCourses, err := base.MapWithErr(courses, toApiCourse)
	if err != nil {
		return
	}

	res = openapi.GetCourses200JSONResponse(apiCourses)

	return
}

// 講義を検索する
// (POST /courses/search)
func (h *impl) PostCoursesSearch(ctx context.Context, request openapi.PostCoursesSearchRequestObject) (res openapi.PostCoursesSearchResponseObject, err error) {
	year, err := shareddomain.ParseAcademicYear(request.Body.Year)
	if err != nil {
		return
	}

	in := timetablemodule.SearchCoursesIn{
		Year:     year,
		Keywords: request.Body.Keywords,
	}

	if request.Body.Codes != nil {
		for _, code := range *request.Body.Codes {
			if strings.HasPrefix(code, "-") {
				in.CodePrefixes.Excluded = append(in.CodePrefixes.Excluded, code[1:])
			} else {
				in.CodePrefixes.Included = append(in.CodePrefixes.Included, code)
			}
		}
	}

	if request.Body.Timetable != nil {
		schedules := fromApiTimetable(*request.Body.Timetable)

		if request.Body.SearchMode == nil || *request.Body.SearchMode == openapi.Cover {
			in.Schedules.PartiallyOverlapped = schedules
		} else {
			in.Schedules.FullyIncluded = schedules
		}
	}

	if request.Body.Limit != nil {
		in.Limit = *request.Body.Limit
	}

	if request.Body.Offset != nil {
		in.Offset = *request.Body.Offset
	}

	courses, err := h.timetableUseCase.SearchCourses(ctx, in)
	if err != nil {
		return
	}

	apiCourses, err := base.MapWithErr(courses, toApiCourse)
	if err != nil {
		return
	}

	res = openapi.PostCoursesSearch200JSONResponse(apiCourses)

	return
}

// 指定した講義を取得する
// (GET /courses/{year}/{code})
func (h *impl) GetCoursesYearCode(ctx context.Context, request openapi.GetCoursesYearCodeRequestObject) (res openapi.GetCoursesYearCodeResponseObject, err error) {
	year, err := shareddomain.ParseAcademicYear(request.Year)
	if err != nil {
		return
	}

	code, err := timetabledomain.ParseCode(request.Code)
	if err != nil {
		return
	}

	courses, err := h.timetableUseCase.GetCoursesByCodes(ctx, year, []timetabledomain.Code{code})
	if err != nil {
		return
	}

	if len(courses) == 0 {
		return nil, apperr.New(sharederr.CodeNotFound, "")
	}

	apiCourse, err := toApiCourse(courses[0])
	if err != nil {
		return
	}

	res = openapi.GetCoursesYearCode200JSONResponse(apiCourse)

	return
}

func fromApiTimetable(q openapi.SearchCourseTimetableQuery) (ret []timetabledomain.Schedule) {
	if q.SpringA != nil {
		schedules := fromApiTimetableDay(*q.SpringA, timetabledomain.ModuleSpringA)
		ret = append(ret, schedules...)
	}

	if q.SpringB != nil {
		schedules := fromApiTimetableDay(*q.SpringB, timetabledomain.ModuleSpringB)
		ret = append(ret, schedules...)
	}

	if q.SpringC != nil {
		schedules := fromApiTimetableDay(*q.SpringC, timetabledomain.ModuleSpringC)
		ret = append(ret, schedules...)
	}

	if q.FallA != nil {
		schedules := fromApiTimetableDay(*q.FallA, timetabledomain.ModuleFallA)
		ret = append(ret, schedules...)
	}

	if q.FallB != nil {
		schedules := fromApiTimetableDay(*q.FallB, timetabledomain.ModuleFallB)
		ret = append(ret, schedules...)
	}

	if q.FallC != nil {
		schedules := fromApiTimetableDay(*q.FallC, timetabledomain.ModuleFallC)
		ret = append(ret, schedules...)
	}

	if q.SummerVacation != nil {
		schedules := fromApiTimetableDay(*q.SummerVacation, timetabledomain.ModuleSummerVacation)
		ret = append(ret, schedules...)
	}

	if q.SpringVacation != nil {
		schedules := fromApiTimetableDay(*q.SpringVacation, timetabledomain.ModuleSpringVacation)
		ret = append(ret, schedules...)
	}

	return
}

func fromApiTimetableDay(q openapi.SearchCourseTimetableQueryDays, module timetabledomain.Module) (ret []timetabledomain.Schedule) {
	if q.Sun != nil {
		periods := fromApiTimetablePeriod(*q.Sun)
		schedules := base.Map(periods, func(period timetabledomain.Period) timetabledomain.Schedule {
			return timetabledomain.Schedule{
				Module: module,
				Day:    timetabledomain.DaySun,
				Period: period,
			}
		})
		ret = append(ret, schedules...)
	}

	if q.Mon != nil {
		periods := fromApiTimetablePeriod(*q.Mon)
		schedules := base.Map(periods, func(period timetabledomain.Period) timetabledomain.Schedule {
			return timetabledomain.Schedule{
				Module: module,
				Day:    timetabledomain.DayMon,
				Period: period,
			}
		})
		ret = append(ret, schedules...)
	}

	if q.Tue != nil {
		periods := fromApiTimetablePeriod(*q.Tue)
		schedules := base.Map(periods, func(period timetabledomain.Period) timetabledomain.Schedule {
			return timetabledomain.Schedule{
				Module: module,
				Day:    timetabledomain.DayTue,
				Period: period,
			}
		})
		ret = append(ret, schedules...)
	}

	if q.Wed != nil {
		periods := fromApiTimetablePeriod(*q.Wed)
		schedules := base.Map(periods, func(period timetabledomain.Period) timetabledomain.Schedule {
			return timetabledomain.Schedule{
				Module: module,
				Day:    timetabledomain.DayWed,
				Period: period,
			}
		})
		ret = append(ret, schedules...)
	}

	if q.Thu != nil {
		periods := fromApiTimetablePeriod(*q.Thu)
		schedules := base.Map(periods, func(period timetabledomain.Period) timetabledomain.Schedule {
			return timetabledomain.Schedule{
				Module: module,
				Day:    timetabledomain.DayThu,
				Period: period,
			}
		})
		ret = append(ret, schedules...)
	}

	if q.Fri != nil {
		periods := fromApiTimetablePeriod(*q.Fri)
		schedules := base.Map(periods, func(period timetabledomain.Period) timetabledomain.Schedule {
			return timetabledomain.Schedule{
				Module: module,
				Day:    timetabledomain.DayFri,
				Period: period,
			}
		})
		ret = append(ret, schedules...)
	}

	if q.Sat != nil {
		periods := fromApiTimetablePeriod(*q.Sat)
		schedules := base.Map(periods, func(period timetabledomain.Period) timetabledomain.Schedule {
			return timetabledomain.Schedule{
				Module: module,
				Day:    timetabledomain.DaySat,
				Period: period,
			}
		})
		ret = append(ret, schedules...)
	}

	if q.Intensive != nil && q.Intensive.N0 != nil && *q.Intensive.N0 {
		ret = append(ret, timetabledomain.Schedule{
			Module: module,
			Day:    timetabledomain.DayIntensive,
		})
	}

	if q.AnyTime != nil && q.AnyTime.N0 != nil && *q.AnyTime.N0 {
		ret = append(ret, timetabledomain.Schedule{
			Module: module,
			Day:    timetabledomain.DayAnyTime,
		})
	}

	if q.Appointment != nil && q.Appointment.N0 != nil && *q.Appointment.N0 {
		ret = append(ret, timetabledomain.Schedule{
			Module: module,
			Day:    timetabledomain.DayAppointment,
		})
	}

	return
}

func fromApiTimetablePeriod(q openapi.SearchCourseTimetableQueryPeriods) (ret []timetabledomain.Period) {
	if q.N1 != nil && *q.N1 {
		ret = append(ret, timetabledomain.Period(1))
	}

	if q.N2 != nil && *q.N2 {
		ret = append(ret, timetabledomain.Period(2))
	}

	if q.N3 != nil && *q.N3 {
		ret = append(ret, timetabledomain.Period(3))
	}

	if q.N4 != nil && *q.N4 {
		ret = append(ret, timetabledomain.Period(4))
	}

	if q.N5 != nil && *q.N5 {
		ret = append(ret, timetabledomain.Period(5))
	}

	if q.N6 != nil && *q.N6 {
		ret = append(ret, timetabledomain.Period(6))
	}

	if q.N7 != nil && *q.N7 {
		ret = append(ret, timetabledomain.Period(7))
	}

	if q.N8 != nil && *q.N8 {
		ret = append(ret, timetabledomain.Period(8))
	}

	return ret
}
