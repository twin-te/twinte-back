package authusecase

import (
	authport "github.com/twin-te/twinte-back/module/auth/port"
)

type Impl struct {
	g authport.Gateway
	r authport.Repository
}

func New(g authport.Gateway, r authport.Repository) *Impl {
	return &Impl{
		g: g,
		r: r,
	}
}
