package liblua

import (
	"fmt"
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

type Fn struct {
	luafn *lua.LFunction
}

func (f *Fn) Run(args ...interface{}) (FnResult, error) {
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
			callback := arg.(func())
			fn := func(*lua.LState) int {
				callback()
				return 0
			}
			luaArgs = append(luaArgs, state.NewFunction(fn))
		default:
			return FnResult{}, fmt.Errorf("not implemented")
		}
	}

	_, err, values := state.Resume(lua.NewState(), f.luafn, luaArgs...)
	if err != nil {
		return FnResult{}, err
	}
	if len(values) == 0 {
		return FnResult{value: lua.LNil}, nil
	}
	return FnResult{value: values[0]}, nil
}
