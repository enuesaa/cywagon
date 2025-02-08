package liblua

import (
	"fmt"
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

type Fn = func (args ...interface{}) (FnResult, error)

func NewFn(luafn *lua.LFunction) Fn {
	return func(args ...interface{}) (FnResult, error) {
		state := lua.NewState()

		luaArgs := []lua.LValue{}
		for _, arg := range args {
			if arg == nil {
				luaArgs = append(luaArgs, lua.LNil)
				continue
			}
			switch reflect.TypeOf(arg).Kind() {
			case reflect.Struct:
				val, err := Marshal(arg)
				if err != nil {
					return FnResult{}, err
				}
				luaArgs = append(luaArgs, val)
			case reflect.String:
				luaArgs = append(luaArgs, lua.LString(arg.(string)))
			case reflect.Int:
				luaArgs = append(luaArgs, lua.LNumber(arg.(int)))
			case reflect.Func:
				callback := arg.(func(t *lua.LTable) *lua.LTable)
				fn := func(s *lua.LState) int {
					table := s.ToTable(1)
					res := callback(table)
					s.Push(res)

					return 1
				}
				luaArgs = append(luaArgs, state.NewFunction(fn))
			default:
				return FnResult{}, fmt.Errorf("not implemented")
			}
		}

		_, err, values := state.Resume(lua.NewState(), luafn, luaArgs...)
		if err != nil {
			return FnResult{}, err
		}
		if len(values) == 0 {
			return FnResult{value: lua.LNil}, nil
		}
		return FnResult{value: values[0]}, nil
	}
}
