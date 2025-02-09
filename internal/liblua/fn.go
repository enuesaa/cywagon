package liblua

import lua "github.com/yuin/gopher-lua"

type Fn func(result interface{}, args ...interface{}) error

func NewFn(luafn *lua.LFunction) Fn {
	return func(result interface{}, args ...interface{}) error {
		var luaArgs []lua.LValue
		for _, arg := range args {
			luaArg, err := Marshal(arg)
			if err != nil {
				return err
			}
			luaArgs = append(luaArgs, luaArg)
		}

		state := lua.NewState()

		_, err, values := state.Resume(state, luafn, luaArgs...)
		if err != nil {
			return err
		}
		if len(values) == 0 {
			return nil
		}

		return Unmarshal(values[0], result)
	}
}
