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

func (r *Runner) Blend(value interface{}) error {
	if err := Inject(r.state, value); err != nil {
		return err
	}
	if err := r.state.DoString(r.code); err != nil {
		return err
	}
	return Eject(r.state, value)
}

func (r *Runner) Run() error {
	return r.state.DoString(r.code)
}

func (r *Runner) GetFunction(name string) Fn {
	luafn := r.state.GetGlobal(name).(*lua.LFunction)

	return Fn{luafn}
}
