package enginectl

import "github.com/enuesaa/cywagon/internal/infra"

func New() Engine {
	return Engine{
		Container: infra.I,
	}
}

type Engine struct {
	infra.Container
}
