package liblua

import (
	"fmt"
	"reflect"

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

func (r *Runner) RunFunction(name string, args... interface{}) (FnResult, error) {
	luaArgs := []lua.LValue{}
	for _, arg := range args {
		if arg == nil {
			luaArgs = append(luaArgs, lua.LNil)
			continue
		}
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Struct:
			luaArgs = append(luaArgs, Parse(arg))
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
			luaArgs = append(luaArgs, r.state.NewFunction(fn))
		default:
			return FnResult{}, fmt.Errorf("not implemented")
		}
	}

	luaFn := r.state.GetGlobal(name).(*lua.LFunction)
	_, err, values := r.state.Resume(lua.NewState(), luaFn, luaArgs...)
	if err != nil {
		return FnResult{}, err
	}
	if len(values) == 0 {
		return FnResult{value: lua.LNil}, nil
	}
	return FnResult{value: values[0]}, nil
}
