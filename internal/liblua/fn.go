package liblua

import lua "github.com/yuin/gopher-lua"

type Fn func(args []interface{}) []interface{}

func NewFn(luafn *lua.LFunction) Fn {
	state := lua.NewState()

	return func(args []interface{}) []interface{} {
		var results []interface{}
		var luaargs []lua.LValue

		_, err, values := state.Resume(lua.NewState(), luafn, luaargs...)
		if err != nil {
			return results
		}
		if len(values) == 0 {
			return results
		}
		return results
	}
}
