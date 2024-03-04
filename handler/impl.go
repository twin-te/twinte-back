package handler

import (
	"net/http"
	"strings"

	restv3 "github.com/twin-te/twinte-back/handler/api/rest/v3"
	apirpc "github.com/twin-te/twinte-back/handler/api/rpc"
	authv3 "github.com/twin-te/twinte-back/handler/auth/v3"
	calendarv1beta "github.com/twin-te/twinte-back/handler/calendar/v1beta"
	announcementmodule "github.com/twin-te/twinte-back/module/announcement"
	authmodule "github.com/twin-te/twinte-back/module/auth"
	donationmodule "github.com/twin-te/twinte-back/module/donation"
	schoolcalendarmodule "github.com/twin-te/twinte-back/module/schoolcalendar"
	timetablemodule "github.com/twin-te/twinte-back/module/timetable"
)

var _ http.Handler = (*impl)(nil)

type impl struct {
	authv3Handler         http.Handler
	calendarv1betaHandler http.Handler
	restv3Handler         http.Handler
	rpcHandler            http.Handler
}

func (h *impl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/api/v3"):
		http.StripPrefix("/api/v3", h.restv3Handler).ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/api"):
		http.StripPrefix("/api", h.rpcHandler).ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/auth/v3"):
		http.StripPrefix("/auth/v3", h.authv3Handler).ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/calendar/v1beta"):
		http.StripPrefix("/calendar/v1beta", h.calendarv1betaHandler).ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}

func New(
	accessController authmodule.AccessController,
	announcementUsecase announcementmodule.UseCase,
	authUseCase authmodule.UseCase,
	donationUseCase donationmodule.UseCase,
	schoolcalendarUseCase schoolcalendarmodule.UseCase,
	timetableUseCase timetablemodule.UseCase,
) *impl {

	authv3Handler := authv3.New(
		accessController,
		authUseCase,
	)

	calendarv1betaHandler := calendarv1beta.New()

	restv3Handler := restv3.New(
		accessController,
		announcementUsecase,
		authUseCase,
		donationUseCase,
		schoolcalendarUseCase,
		timetableUseCase,
	)

	rpcHandler := apirpc.New(
		accessController,
		announcementUsecase,
		authUseCase,
		schoolcalendarUseCase,
		timetableUseCase,
	)

	h := &impl{
		authv3Handler:         authv3Handler,
		calendarv1betaHandler: calendarv1betaHandler,
		restv3Handler:         restv3Handler,
		rpcHandler:            rpcHandler,
	}

	return h
}
