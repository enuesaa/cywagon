package liblua

import (
	"github.com/enuesaa/cywagon/internal/infra"
	lua "github.com/yuin/gopher-lua"
)

func New(container infra.Container) Runner {
	runner := Runner{
		Container: container,
		state:     lua.NewState(),
	}
	return runner
}

type Runner struct {
	infra.Container
	state *lua.LState
}

func (r *Runner) Run(code string) error {
	return r.state.DoString(code)
}
