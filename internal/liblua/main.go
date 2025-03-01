package liblua

import lua "github.com/yuin/gopher-lua"

func New() Runner {
	runner := Runner{
		state:     lua.NewState(),
	}
	return runner
}

type Runner struct {
	state *lua.LState
}

func (r *Runner) Run(code string) error {
	return r.state.DoString(code)
}
