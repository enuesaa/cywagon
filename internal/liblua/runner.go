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

func (r *Runner) Run() error {
	entry := lua.LTable{}
	r.state.SetField(&entry, "cmd", lua.LString("a"))
	r.state.SetField(&entry, "workdir", lua.LString("a"))
	r.state.SetField(&entry, "waitForHealthy", lua.LNumber(60))
	r.state.SetGlobal("entry", &entry)

	healthCheck := lua.LTable{}
	r.state.SetField(&healthCheck, "protocol", lua.LString("a"))
	r.state.SetField(&healthCheck, "method", lua.LString("a"))
	r.state.SetField(&healthCheck, "path", lua.LNumber(60))
	r.state.SetGlobal("healthCheck", &healthCheck)

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
