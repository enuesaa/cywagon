package liblua

import lua "github.com/yuin/gopher-lua"

func NewRunner(code string) Runner {
	return Runner{
		code:  code,
		state: lua.NewState(),
	}
}

type Runner struct {
	code  string
	state *lua.LState
}

func (r *Runner) Inject(value interface{}) error {
	return Inject(r.state, value)
}

func (r *Runner) Run() error {
	return r.state.DoString(r.code)
}

func (r *Runner) Eject(value interface{}) error {
	return Eject(r.state, value)
}
