package enginectl

import "github.com/enuesaa/cywagon/internal/infra"

func New(ctn infra.Container) Engine {
	return Engine{
		ctn: ctn,
	}
}

type Engine struct {
	ctn infra.Container
}
