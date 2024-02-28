package restv3

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	"github.com/twin-te/twinte-back/handler/api/rest/v3/openapi"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	timetablemodule "github.com/twin-te/twinte-back/module/timetable"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

func toApiRegisteredCourse(registeredCourse *timetabledomain.RegisteredCourse, idToCourse map[idtype.CourseID]*timetabledomain.Course) (ret openapi.RegisteredCourse, err error) {
	ret = openapi.RegisteredCourse{
		Absence:    int(registeredCourse.Absence),
		Attendance: int(registeredCourse.Attendance),
		Id:         toApiUUID(registeredCourse.ID),
		Instructor: registeredCourse.Instructors,
		Late:       int(registeredCourse.Late),
		Memo:       registeredCourse.Memo,
		Tags: base.Map(registeredCourse.TagIDs, func(tagID idtype.TagID) openapi.TagIdOnly {
			return openapi.TagIdOnly{
				Id: toApiUUID(tagID),
			}
		}),
		UserId: toApiUUID(registeredCourse.UserID),
		Year:   registeredCourse.Year.Int(),
	}

	if registeredCourse.CourseID != nil {
		course, err := toApiCourse(idToCourse[*registeredCourse.CourseID])
		if err != nil {
			return openapi.RegisteredCourse{}, err
		}
		ret.Course = &course
	}

	if registeredCourse.Credit != nil {
		ret.Credit, err = base.ToPtrWithErr(toApiCredit(*registeredCourse.Credit))
		if err != nil {
			return
		}
	}

	if registeredCourse.Methods != nil {
		ret.Methods, err = base.ToPtrWithErr(base.MapWithErr(*registeredCourse.Methods, toApiCourseMethod))
		if err != nil {
			return
		}
	}

	if registeredCourse.Name != nil {
		ret.Name = lo.ToPtr(registeredCourse.Name.String())
	}

	if registeredCourse.Schedules != nil {
		ret.Schedules, err = base.ToPtrWithErr(base.MapWithErr(*registeredCourse.Schedules, toApiCourseSchedule))
		if err != nil {
			return
		}
	}

	return
}

// 登録済みの講義を返す
// (GET /registered-courses)
func (h *impl) GetRegisteredCourses(ctx context.Context, request openapi.GetRegisteredCoursesRequestObject) (res openapi.GetRegisteredCoursesResponseObject, err error) {
	year, err := shareddomain.ParseAcademicYear(request.Params.Year)
	if err != nil {
		return
	}

	registeredCourses, err := h.timetableUseCase.GetRegisteredCourses(ctx, &year)
	if err != nil {
		return
	}

	apiRegisteredCourses, err := h.getApiRegisteredCourses(ctx, registeredCourses)
	if err != nil {
		return
	}

	res = openapi.GetRegisteredCourses200JSONResponse(apiRegisteredCourses)

	return
}

func (h *impl) postRegisteredCourses0(ctx context.Context, reqBody openapi.PostRegisteredCoursesJSONBody0) (apiRegisteredCourse openapi.RegisteredCourse, err error) {
	year, err := shareddomain.ParseAcademicYear(reqBody.Year)
	if err != nil {
		return
	}

	code, err := timetabledomain.ParseCode(reqBody.Code)
	if err != nil {
		return
	}

	registeredCourses, err := h.timetableUseCase.CreateRegisteredCoursesByCodes(ctx, year, []timetabledomain.Code{code})
	if err != nil {
		return
	}

	apiRegisteredCourses, err := h.getApiRegisteredCourses(ctx, registeredCourses)
	if err != nil {
		return
	}

	return apiRegisteredCourses[0], nil
}

func (h *impl) postRegisteredCourses1(ctx context.Context, reqBody openapi.PostRegisteredCoursesJSONBody1) ([]openapi.RegisteredCourse, error) {
	yearToCodes := make(map[shareddomain.AcademicYear][]timetabledomain.Code)

	for i := range reqBody {
		year, err := shareddomain.ParseAcademicYear(reqBody[i].Year)
		if err != nil {
			return nil, err
		}

		code, err := timetabledomain.ParseCode(reqBody[i].Code)
		if err != nil {
			return nil, err
		}

		yearToCodes[year] = append(yearToCodes[year], code)
	}

	var registeredCourses []*timetabledomain.RegisteredCourse

	for year, codes := range yearToCodes {
		rcs, err := h.timetableUseCase.CreateRegisteredCoursesByCodes(ctx, year, codes)
		if err != nil {
			return nil, err
		}

		registeredCourses = append(registeredCourses, rcs...)
	}

	return h.getApiRegisteredCourses(ctx, registeredCourses)
}

func (h *impl) postRegisteredCourses2(ctx context.Context, reqBody openapi.PostRegisteredCoursesJSONBody2) (apiRegisteredCourse openapi.RegisteredCourse, err error) {
	in := timetablemodule.CreateRegisteredCourseManuallyIn{
		Instructors: reqBody.Instructor,
	}

	in.Year, err = shareddomain.ParseAcademicYear(reqBody.Year)
	if err != nil {
		return
	}

	in.Name, err = timetabledomain.ParseName(reqBody.Name)
	if err != nil {
		return
	}

	in.Credit, err = fromApiCredit(reqBody.Credit)
	if err != nil {
		return
	}

	in.Methods, err = base.MapWithErr(reqBody.Methods, fromApiCourseMethod)
	if err != nil {
		return
	}

	in.Schedules, err = base.MapWithErr(reqBody.Schedules, fromApiCourseSchedule)
	if err != nil {
		return
	}

	registeredCourse, err := h.timetableUseCase.CreateRegisteredCourseManually(ctx, in)
	if err != nil {
		return
	}

	apiRegisteredCourses, err := h.getApiRegisteredCourses(ctx, []*timetabledomain.RegisteredCourse{registeredCourse})
	if err != nil {
		return
	}

	return apiRegisteredCourses[0], nil
}

// 講義を登録する
// (POST /registered-courses)
func (h *impl) PostRegisteredCourses(ctx context.Context, request openapi.PostRegisteredCoursesRequestObject) (openapi.PostRegisteredCoursesResponseObject, error) {
	if reqBody, err := openapi.ToPostRegisteredCoursesJSONBody0(request.Body); err != nil {
		apiRegisteredCourse, err := h.postRegisteredCourses0(ctx, reqBody)
		if err != nil {
			return nil, err
		}
		return openapi.FromRegisteredCourse(apiRegisteredCourse)
	}

	if reqBody, err := openapi.ToPostRegisteredCoursesJSONBody1(request.Body); err != nil {
		apiRegisteredCourses, err := h.postRegisteredCourses1(ctx, reqBody)
		if err != nil {
			return nil, err
		}
		return openapi.FromRegisteredCourses(apiRegisteredCourses)
	}

	if reqBody, err := openapi.ToPostRegisteredCoursesJSONBody2(request.Body); err != nil {
		apiRegisteredCourse, err := h.postRegisteredCourses2(ctx, reqBody)
		if err != nil {
			return nil, err
		}
		return openapi.FromRegisteredCourse(apiRegisteredCourse)
	}

	return nil, echo.ErrBadRequest
}

// 指定した登録された講義を削除する
// (DELETE /registered-courses/{id})
func (h *impl) DeleteRegisteredCoursesId(ctx context.Context, request openapi.DeleteRegisteredCoursesIdRequestObject) (res openapi.DeleteRegisteredCoursesIdResponseObject, err error) {
	id, err := idtype.ParseRegisteredCourseID(request.Id.String())
	if err != nil {
		return
	}

	err = h.timetableUseCase.DeleteRegisteredCourse(ctx, id)
	if err != nil {
		return
	}

	res = openapi.DeleteRegisteredCoursesId204Response{}

	return
}

// 指定した登録された講義を取得する
// (GET /registered-courses/{id})
func (h *impl) GetRegisteredCoursesId(ctx context.Context, request openapi.GetRegisteredCoursesIdRequestObject) (res openapi.GetRegisteredCoursesIdResponseObject, err error) {
	id, err := idtype.ParseRegisteredCourseID(request.Id.String())
	if err != nil {
		return
	}

	registeredCourse, err := h.timetableUseCase.GetRegisteredCourseByID(ctx, id)
	if err != nil {
		return
	}

	apiRegisteredCourses, err := h.getApiRegisteredCourses(ctx, []*timetabledomain.RegisteredCourse{registeredCourse})
	if err != nil {
		return
	}

	res = openapi.GetRegisteredCoursesId200JSONResponse(apiRegisteredCourses[0])

	return
}

// 指定した登録された講義を更新する
// (PUT /registered-courses/{id})
func (h *impl) PutRegisteredCoursesId(ctx context.Context, request openapi.PutRegisteredCoursesIdRequestObject) (res openapi.PutRegisteredCoursesIdResponseObject, err error) {
	id, err := idtype.ParseRegisteredCourseID(request.Id.String())
	if err != nil {
		return
	}

	in := timetablemodule.UpdateRegisteredCourseIn{
		ID:          id,
		Instructors: request.Body.Instructor,
		Memo:        &request.Body.Memo,
	}

	if request.Body.Name != nil {
		in.Name, err = base.ToPtrWithErr(timetabledomain.ParseName(*request.Body.Name))
		if err != nil {
			return
		}
	}

	if request.Body.Credit != nil {
		in.Credit, err = base.ToPtrWithErr(fromApiCredit(*request.Body.Credit))
		if err != nil {
			return
		}
	}

	if request.Body.Methods != nil {
		in.Methods, err = base.ToPtrWithErr(base.MapWithErr(*request.Body.Methods, fromApiCourseMethod))
		if err != nil {
			return
		}
	}

	if request.Body.Schedules != nil {
		in.Schedules, err = base.ToPtrWithErr(base.MapWithErr(*request.Body.Schedules, fromApiCourseSchedule))
		if err != nil {
			return
		}
	}

	in.Attendance, err = base.ToPtrWithErr(timetabledomain.ParseAttendance(request.Body.Attendance))
	if err != nil {
		return
	}

	in.Absence, err = base.ToPtrWithErr(timetabledomain.ParseAbsence(request.Body.Absence))
	if err != nil {
		return
	}

	in.Late, err = base.ToPtrWithErr(timetabledomain.ParseLate(request.Body.Late))
	if err != nil {
		return
	}

	in.TagIDs, err = base.ToPtrWithErr(
		base.MapWithErr(request.Body.Tags, func(tagIdOnly openapi.TagIdOnly) (idtype.TagID, error) {
			return idtype.ParseTagID(tagIdOnly.Id.String())
		}),
	)
	if err != nil {
		return
	}

	registeredCourse, err := h.timetableUseCase.UpdateRegisteredCourse(ctx, in)
	if err != nil {
		return
	}

	apiRegisteredCourses, err := h.getApiRegisteredCourses(ctx, []*timetabledomain.RegisteredCourse{registeredCourse})
	if err != nil {
		return
	}

	res = openapi.PutRegisteredCoursesId200JSONResponse(apiRegisteredCourses[0])

	return
}

func (h *impl) getApiRegisteredCourses(ctx context.Context, registeredCourses []*timetabledomain.RegisteredCourse) ([]openapi.RegisteredCourse, error) {
	courseIDs := make([]idtype.CourseID, 0, len(registeredCourses))
	for _, registeredCourse := range registeredCourses {
		if registeredCourse.CourseID != nil {
			courseIDs = append(courseIDs, *registeredCourse.CourseID)
		}
	}

	courses, err := h.timetableUseCase.GetCoursesByIDs(ctx, courseIDs)
	if err != nil {
		return nil, err
	}
	if len(courseIDs) != len(courses) {
		return nil, fmt.Errorf("not found courses in getApiRegisteredCourses %+v", courseIDs)
	}

	idToCourse := lo.SliceToMap(courses, func(course *timetabledomain.Course) (idtype.CourseID, *timetabledomain.Course) {
		return course.ID, course
	})

	return base.MapWithArgAndErr(registeredCourses, idToCourse, toApiRegisteredCourse)
}
