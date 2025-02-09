package liblua

import (
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

type Fn = func(args []interface{}) []interface{}

func NewFn(luafn *lua.LFunction) Fn {
	state := lua.NewState()

	return func(args []interface{}) []interface{} {
		var results []interface{}
		var luaargs []lua.LValue

		for _, arg := range args {
			argType := reflect.TypeOf(arg)

			switch argType.Kind() {
			case reflect.Struct:
				val, err := Marshal(arg)
				if err != nil {
					return results
				}
				luaargs = append(luaargs, val)
			case reflect.Func:
				callback := arg.(func(interface{}) interface{})

				fnargs := []lua.LValue{}
				for i := range argType.NumIn() {
					val, _ := Marshal(argType.In(i))
					fnargs = append(fnargs, val)
				}

				fn := func(s *lua.LState) int {
					table := s.ToTable(1)
					res := callback(table)
					luares, _ := Marshal(res)
					s.Push(luares)
					return 1
				}
				luaargs = append(luaargs, state.NewFunction(fn))
			default:
				return results
			}
		}

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
