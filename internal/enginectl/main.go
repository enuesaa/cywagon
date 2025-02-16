package enginectl

import "github.com/enuesaa/cywagon/internal/infra"

func New() Engine {
	engine := Engine{
		Container: infra.Default,
	}
	return engine
}

type Engine struct {
	infra.Container
}
