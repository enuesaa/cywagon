package enginectl

import (
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/libsock"
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
)

type EngineInterface interface {
	Serve(config model.Config, workdir string) error
	StartListenSock() error
	Close() error
}

func New() *Engine {
	engine := Engine{
		Container: infra.Default,
		Server:    libserve.New(),
		ConfSrv:   service.NewConfSrv(),
		Sock:      libsock.New(),
	}
	return &engine
}

type Engine struct {
	infra.Container

	Server  libserve.Server
	ConfSrv service.ConfSrvInterface
	Sock    libsock.Sock
}
