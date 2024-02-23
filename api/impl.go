package api

import (
	"net/http"
	"strings"

	apirpc "github.com/twin-te/twinte-back/api/rpc"
	announcementmodule "github.com/twin-te/twinte-back/module/announcement"
	authmodule "github.com/twin-te/twinte-back/module/auth"
	schoolcalendarmodule "github.com/twin-te/twinte-back/module/schoolcalendar"
	timetablemodule "github.com/twin-te/twinte-back/module/timetable"
)

var _ http.Handler = (*impl)(nil)

type impl struct {
	RESTHandler http.Handler
	RPCHandler  http.Handler
}

func (h *impl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/rest"):
		http.StripPrefix("/rest", h.RESTHandler)
	case strings.HasPrefix(r.URL.Path, "/rpc"):
		http.StripPrefix("/rpc", h.RPCHandler)
	default:
		http.NotFound(w, r)
	}
}

func New(
	announcementUsecase announcementmodule.UseCase,
	authUseCase authmodule.UseCase,
	schoolcalendarUseCase schoolcalendarmodule.UseCase,
	timetableUseCase timetablemodule.UseCase,
) *impl {
	rpcHandler := apirpc.New(
		announcementUsecase,
		authUseCase,
		schoolcalendarUseCase,
		timetableUseCase,
	)

	h := &impl{
		RPCHandler: rpcHandler,
	}

	return h
}
