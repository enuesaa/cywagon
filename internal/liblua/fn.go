package liblua

import (
	"fmt"
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

type FnI = func(args []interface{}) []interface{}


type Fn = func(args ...interface{}) (*lua.LTable, error)

func NewFn(luafn *lua.LFunction) FnI {
	return func(args []interface{}) []interface{} {
		results := make([]interface{}, 0)
		state := lua.NewState()

		var luaArgs []lua.LValue
		for _, arg := range args {
			if arg == nil {
				luaArgs = append(luaArgs, lua.LNil)
				continue
			}
			switch reflect.TypeOf(arg).Kind() {
			case reflect.Struct:
				val, err := Marshal(arg)
				if err != nil {
					return results
				}
				luaArgs = append(luaArgs, val)
			case reflect.Func:
				callback := arg.(func(interface{}) interface{})
				fn := func(s *lua.LState) int {
					table := s.ToTable(1)
					res := callback(table)
					luares, _ := Marshal(res)
					s.Push(luares)
					return 1
				}
				luaArgs = append(luaArgs, state.NewFunction(fn))
			default:
				return results
			}
		}

		_, err, values := state.Resume(lua.NewState(), luafn, luaArgs...)
		if err != nil {
			return results
		}
		if len(values) == 0 {
			return results
		}
		return results
	}
}
