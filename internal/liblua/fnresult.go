package liblua

import lua "github.com/yuin/gopher-lua"

type FnResult struct {
	value lua.LValue
}

func (r *FnResult) GetInt(name string) int {
	state := lua.NewState()
	target := state.GetField(r.value, name).(lua.LNumber)

	return int(target)
}
