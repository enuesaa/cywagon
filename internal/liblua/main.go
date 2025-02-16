package liblua

import (
	"github.com/enuesaa/cywagon/internal/infra"
	lua "github.com/yuin/gopher-lua"
)

func New() Runner {
	return Runner{
		Container: infra.Default,
		state: lua.NewState(),
	}
}

type Runner struct {
	infra.Container
	state *lua.LState
}

func (r *Runner) Run(code string) error {
	return r.state.DoString(code)
}
