package enginectl

import (
	"io/fs"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
)

type EngineInterface interface {
	Load(config model.Config, workdir string) error
	Serve() error
	Close() error
}

func New() *Engine {
	engine := Engine{
		Container: infra.Default,
		Server:    libserve.New(),
		ConfSrv:   service.NewConfSrv(),
		Log:       service.NewLogSrv(),

		sitemap:  make(map[string]model.Site),
		distmap:  make(map[string]fs.FS),
		logicmap: make(map[string]model.Logic),
	}
	return &engine
}

type Engine struct {
	infra.Container

	Server  libserve.Server
	ConfSrv service.ConfSrvInterface
	Log     service.LogSrvInterface

	sitemap  map[string]model.Site
	distmap  map[string]fs.FS
	logicmap map[string]model.Logic
}
