package timetableusecase

import (
	authmodule "github.com/twin-te/twinte-back/module/auth"
	timetableport "github.com/twin-te/twinte-back/module/timetable/port"
)

type Impl struct {
	a authmodule.Authorizer
	r timetableport.Repository
}

func New(a authmodule.Authorizer, r timetableport.Repository) *Impl {
	return &Impl{
		a: a,
		r: r,
	}
}
