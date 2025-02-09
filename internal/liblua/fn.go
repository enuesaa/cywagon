package liblua

import lua "github.com/yuin/gopher-lua"

type Fn func(arg interface{}, result interface{}) error

func NewFn(luafn *lua.LFunction) Fn {
	return func(arg interface{}, result interface{}) error {
		luaArg, err := Marshal(arg)
		if err != nil {
			return err
		}

		state := lua.NewState()

		_, err, values := state.Resume(state, luafn, luaArg)
		if err != nil {
			return err
		}
		if len(values) == 0 {
			return nil
		}

		return Unmarshal(values[0], result)
	}
}
