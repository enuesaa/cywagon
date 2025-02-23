package enginectl

import (
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libserve"
)

func New(container infra.Container) Engine {
	engine := Engine{
		Container: container,
		Server: libserve.Server{
			Container: container,
		},
	}
	return engine
}

type Engine struct {
	infra.Container

	Server libserve.Server
}
