package liblua

import (
	lua "github.com/yuin/gopher-lua"
)

func NewRunner(code string) Runner {
	return Runner{
		code: code,
		state: lua.NewState(),
	}
}

type Runner struct {
	code string
	state *lua.LState

}

func (r *Runner) S() *lua.LState {
	return r.state
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

func (r *Runner) RunFunction(name string, args []lua.LValue) ([]lua.LValue, error) {
	fn := r.state.GetGlobal(name).(*lua.LFunction)
	_, err, values := r.state.Resume(lua.NewState(), fn, args...)
	if err != nil {
		return values, err
	}
	return values, nil
}
