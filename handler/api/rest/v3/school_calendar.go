package restv3

import (
	"context"
	"fmt"
	"time"

	"github.com/twin-te/twinte-back/base"
	"github.com/twin-te/twinte-back/handler/api/rest/v3/openapi"
	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func toApiCourseDay(day time.Weekday) (openapi.CourseDay, error) {
	switch day {
	case time.Sunday:
		return openapi.CourseDaySun, nil
	case time.Monday:
		return openapi.CourseDayMon, nil
	case time.Tuesday:
		return openapi.CourseDayTue, nil
	case time.Wednesday:
		return openapi.CourseDayWed, nil
	case time.Thursday:
		return openapi.CourseDayThu, nil
	case time.Friday:
		return openapi.CourseDayFri, nil
	case time.Saturday:
		return openapi.CourseDaySat, nil
	}
	return "", fmt.Errorf("invalid %#v", day)
}

func toApiSchoolCalendarEventEventType(eventType schoolcalendardomain.EventType) (openapi.SchoolCalendarEventEventType, error) {
	switch eventType {
	case schoolcalendardomain.EventTypeHoliday:
		return openapi.Holiday, nil
	case schoolcalendardomain.EventTypePublicHoliday:
		return openapi.PublicHoliday, nil
	case schoolcalendardomain.EventTypeExam:
		return openapi.Exam, nil
	case schoolcalendardomain.EventTypeSubstituteDay:
		return openapi.SubstituteDay, nil
	case schoolcalendardomain.EventTypeOther:
		return openapi.Other, nil
	}
	return "", fmt.Errorf("invalid %#v", eventType)
}

func toApiSchoolCalendarModuleModule(module schoolcalendardomain.Module) (openapi.SchoolCalendarModuleModule, error) {
	switch module {
	case schoolcalendardomain.ModuleSpringA:
		return openapi.SpringA, nil
	case schoolcalendardomain.ModuleSpringB:
		return openapi.SpringB, nil
	case schoolcalendardomain.ModuleSpringC:
		return openapi.SpringC, nil
	case schoolcalendardomain.ModuleSummerVacation:
		return openapi.SummerVacation, nil
	case schoolcalendardomain.ModuleFallA:
		return openapi.FallA, nil
	case schoolcalendardomain.ModuleFallB:
		return openapi.FallB, nil
	case schoolcalendardomain.ModuleWinterVacation:
		// TODO: create api winter vacation
		return openapi.FallC, nil
	case schoolcalendardomain.ModuleFallC:
		return openapi.FallC, nil
	case schoolcalendardomain.ModuleSpringVacation:
		return openapi.SpringVacation, nil
	}
	return "", fmt.Errorf("invalid %#v", module)
}

func toApiSchoolCalendarEvent(event *schoolcalendardomain.Event) (ret openapi.SchoolCalendarEvent, err error) {
	ret.Date = openapi_types.Date{Time: event.Date.In(time.Local)}
	ret.Description = event.Description

	ret.EventType, err = toApiSchoolCalendarEventEventType(event.Type)
	if err != nil {
		return
	}

	if event.ChangeTo != nil {
		ret.ChangeTo, err = base.ToPtrWithErr(toApiCourseDay(*event.ChangeTo))
		if err != nil {
			return
		}
	}

	return
}

func toApiModuleDetail(moduleDetail *schoolcalendardomain.ModuleDetail) (ret openapi.SchoolCalendarModule, err error) {
	ret.Year = int(moduleDetail.Year)
	ret.Start = openapi_types.Date{Time: moduleDetail.Start.In(time.Local)}
	ret.End = openapi_types.Date{Time: moduleDetail.End.In(time.Local)}
	ret.Module, err = toApiSchoolCalendarModuleModule(moduleDetail.Module)
	return
}

// 学年暦のイベントを取得する
// (GET /school-calendar/events)
func (h *impl) GetSchoolCalendarEvents(ctx context.Context, request openapi.GetSchoolCalendarEventsRequestObject) (res openapi.GetSchoolCalendarEventsResponseObject, err error) {
	year, err := shareddomain.ParseAcademicYear(request.Params.Year)
	if err != nil {
		return
	}

	events, err := h.schoolCalendarUseCase.GetEvents(ctx, year)
	if err != nil {
		return
	}

	apiEvents, err := base.MapWithErr(events, toApiSchoolCalendarEvent)
	if err != nil {
		return
	}

	res = openapi.GetSchoolCalendarEvents200JSONResponse(apiEvents)

	return
}

// 学年暦のモジュール期間を取得する
// (GET /school-calendar/modules)
func (h *impl) GetSchoolCalendarModules(ctx context.Context, request openapi.GetSchoolCalendarModulesRequestObject) (res openapi.GetSchoolCalendarModulesResponseObject, err error) {
	year, err := shareddomain.ParseAcademicYear(request.Params.Year)
	if err != nil {
		return
	}

	moduleDetails, err := h.schoolCalendarUseCase.GetModuleDetails(ctx, year)
	if err != nil {
		return
	}

	apiModuleDetails, err := base.MapWithErr(moduleDetails, toApiModuleDetail)
	if err != nil {
		return
	}

	res = openapi.GetSchoolCalendarModules200JSONResponse(apiModuleDetails)

	return
}
