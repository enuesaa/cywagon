package enginectl

import (
	"github.com/enuesaa/cywagon/internal/infra"
)

func New() Engine {
	return Engine{
		Container: infra.Default,
	}
}

type Engine struct {
	infra.Container
}
