package handle

import (
	"github.com/enuesaa/cywagon/internal/enginectl"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service"
)

func New(container infra.Container) Handler {
	return Handler{
		Container: container,
		Engine: enginectl.New(container),
		ConfSrv: service.NewConfService(container),
	}
}

type Handler struct {
	infra.Container
	Engine enginectl.EngineCtl
	ConfSrv service.ConfService
}
