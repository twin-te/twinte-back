package restv3

import (
	"context"

	"cloud.google.com/go/civil"
	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	"github.com/twin-te/twinte-back/handler/api/rest/v3/openapi"
	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
)

// 指定された日付の日程＆時間割情報を取得する
// (GET /timetable/{date})
func (h *impl) GetTimetableDate(ctx context.Context, request openapi.GetTimetableDateRequestObject) (res openapi.GetTimetableDateResponseObject, err error) {
	date := civil.DateOf(request.Date.Time)

	year, err := shareddomain.NewAcademicYear(date.Year, date.Month)
	if err != nil {
		return
	}

	events, err := h.schoolCalendarUseCase.GetEventsByDate(ctx, date)
	if err != nil {
		return
	}

	apiEvents, err := base.MapWithErr(events, toApiSchoolCalendarEvent)
	if err != nil {
		return
	}

	moduleDetails, err := h.schoolCalendarUseCase.GetModuleDetails(ctx, year)
	if err != nil {
		return
	}

	moduleDetail, ok := lo.Find(moduleDetails, func(item *schoolcalendardomain.ModuleDetail) bool {
		return !item.Start.After(date) && !item.End.Before(date)
	})

	var apiModuleDetail *openapi.SchoolCalendarModule

	if ok {
		apiModuleDetail, err = base.ToPtrWithErr(toApiModuleDetail(moduleDetail))
		if err != nil {
			return
		}
	}

	registeredCourses, err := h.timetableUseCase.GetRegisteredCourses(ctx, &year)
	if err != nil {
		return
	}

	apiRegisteredCourses, err := h.getApiRegisteredCourses(ctx, registeredCourses)
	if err != nil {
		return
	}

	res = openapi.GetTimetableDate200JSONResponse{
		Courses: apiRegisteredCourses,
		Date:    request.Date,
		Events:  apiEvents,
		Module:  apiModuleDetail,
	}

	return
}
