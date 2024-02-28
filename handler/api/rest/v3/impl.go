package restv3

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/twin-te/twinte-back/handler/api/rest/v3/openapi"
	"github.com/twin-te/twinte-back/handler/common/middleware"
	announcementmodule "github.com/twin-te/twinte-back/module/announcement"
	authmodule "github.com/twin-te/twinte-back/module/auth"
	donationmodule "github.com/twin-te/twinte-back/module/donation"
	schoolcalendarmodule "github.com/twin-te/twinte-back/module/schoolcalendar"
	timetablemodule "github.com/twin-te/twinte-back/module/timetable"
)

var _ openapi.StrictServerInterface = (*impl)(nil)

type impl struct {
	accessController authmodule.AccessController

	announcementUseCase   announcementmodule.UseCase
	authUseCase           authmodule.UseCase
	donationUseCase       donationmodule.UseCase
	schoolCalendarUseCase schoolcalendarmodule.UseCase
	timetableUseCase      timetablemodule.UseCase
}

func New(
	accessController authmodule.AccessController,
	announcementUseCase announcementmodule.UseCase,
	authUseCase authmodule.UseCase,
	donationUseCase donationmodule.UseCase,
	schoolCalendarUseCase schoolcalendarmodule.UseCase,
	timetableUseCase timetablemodule.UseCase,
) *echo.Echo {
	e := echo.New()

	e.Use(
		echomiddleware.Recover(),
		echomiddleware.Logger(),
		middleware.NewEchoErrorHandler(),
		middleware.NewEchoWithActor(accessController),
	)

	handler := &impl{
		accessController:      accessController,
		announcementUseCase:   announcementUseCase,
		authUseCase:           authUseCase,
		donationUseCase:       donationUseCase,
		schoolCalendarUseCase: schoolCalendarUseCase,
		timetableUseCase:      timetableUseCase,
	}
	strictHandler := openapi.NewStrictHandler(handler, nil)
	openapi.RegisterHandlers(e, strictHandler)

	return e
}
