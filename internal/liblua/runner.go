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

func (r *Runner) SetGlobal(name string, value interface{}) {
	r.state.SetGlobal(name, Marshal(value))
}

func (r *Runner) Run() error {
	return r.state.DoString(r.code)
}

func (r *Runner) GetString(name string) string {
	value := r.state.GetGlobal(name).(lua.LString)

	return value.String()
}

func (r *Runner) GetInt(name string) int {
	value := r.state.GetGlobal(name).(lua.LNumber)

	return int(value)
}

func (r *Runner) GetFunction(name string) Fn {
	luafn := r.state.GetGlobal(name).(*lua.LFunction)

	return Fn{luafn}
}
