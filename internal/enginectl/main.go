package enginectl

import (
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libfetch"
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service"
)

func New(container infra.Container) Engine {
	engine := Engine{
		Container: container,
		Server: libserve.New(container),
		ConfSrv: service.NewConfService(container),
		Fetcher: libfetch.New(container),
	}
	return engine
}

type Engine struct {
	infra.Container

	Server libserve.Server
	ConfSrv service.ConfService
	Fetcher libfetch.Fetcher
}
