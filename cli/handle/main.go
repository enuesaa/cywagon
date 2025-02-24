package handle

import (
	"github.com/enuesaa/cywagon/internal/enginectl"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service"
)

func New() Handler {
	return Handler{
		Container: infra.Default,
		Engine: enginectl.New(),
		ConfSrv: service.NewConfService(),
	}
}

type Handler struct {
	infra.Container
	Engine enginectl.EngineCtl
	ConfSrv service.ConfServicer
}
