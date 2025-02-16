package enginectl

import "github.com/enuesaa/cywagon/internal/infra"

func New(opts ...func(*Engine)) Engine {
	engine := Engine{
		Container: infra.Default,
	}
	for _, opt := range opts {
		opt(&engine)
	}
	return engine
}

type Engine struct {
	infra.Container
}
