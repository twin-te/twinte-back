package restv3

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/twin-te/twinte-back/api/rest/v3/openapi"
	"github.com/twin-te/twinte-back/apperr"
	announcementmodule "github.com/twin-te/twinte-back/module/announcement"
	authmodule "github.com/twin-te/twinte-back/module/auth"
	donationmodule "github.com/twin-te/twinte-back/module/donation"
	schoolcalendarmodule "github.com/twin-te/twinte-back/module/schoolcalendar"
	sharederr "github.com/twin-te/twinte-back/module/shared/err"
	timetablemodule "github.com/twin-te/twinte-back/module/timetable"
)

var _ openapi.StrictServerInterface = (*impl)(nil)

type impl struct {
	announcementUseCase   announcementmodule.UseCase
	authUseCase           authmodule.UseCase
	donationUseCase       donationmodule.UseCase
	schoolCalendarUseCase schoolcalendarmodule.UseCase
	timetableUseCase      timetablemodule.UseCase
}

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		if aerr, ok := apperr.As(err); ok {
			switch aerr.Code {
			case sharederr.CodeInvalidArgument:
				return echo.NewHTTPError(http.StatusBadRequest, aerr.Message)
			case sharederr.CodeUnauthenticated:
				return echo.NewHTTPError(http.StatusUnauthorized, aerr.Message)
			case sharederr.CodeNotFound:
				return echo.NewHTTPError(http.StatusNotFound, aerr.Message)
			}
		}

		return echo.NewHTTPError(http.StatusInternalServerError)
	}
}

func New(
	announcementUseCase announcementmodule.UseCase,
	authUseCase authmodule.UseCase,
	donationUseCase donationmodule.UseCase,
	schoolCalendarUseCase schoolcalendarmodule.UseCase,
	timetableUseCase timetablemodule.UseCase,
) http.Handler {
	e := echo.New()
	e.Use(ErrorHandler)
	handler := &impl{
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
