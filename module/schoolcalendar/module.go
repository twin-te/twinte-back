package schoolcalendarmodule

import (
	"context"

	"cloud.google.com/go/civil"
	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
)

// UseCase represents application specific business rules.
//
// The error codes for authentication and authorization failures are not stated explicitly.
type UseCase interface {
	// GetEvents returns the events specified by the given year.
	//
	// [Authentication] not required
	GetEvents(ctx context.Context, year shareddomain.AcademicYear) ([]*schoolcalendardomain.Event, error)

	// GetEventsByDate returns the events specified by the given date.
	//
	// [Authentication] not required
	GetEventsByDate(ctx context.Context, date civil.Date) ([]*schoolcalendardomain.Event, error)

	// GetModuleDetails returns the module details specified by the given year.
	//
	// [Authentication] not required
	GetModuleDetails(ctx context.Context, year shareddomain.AcademicYear) ([]*schoolcalendardomain.ModuleDetail, error)
}
