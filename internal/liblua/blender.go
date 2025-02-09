package liblua

import (
	"fmt"
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

func Inject(state *lua.LState, from interface{}) error {
	fromType := reflect.TypeOf(from)
	fromReal := reflect.ValueOf(from)
	if fromType.Kind() != reflect.Struct {
		return fmt.Errorf("unsupported value supplied")
	}

	for i := range fromType.NumField() {
		field := fromType.Field(i)
		value := fromReal.Field(i).Interface()

		name, err := extarctLuaTagValue(field.Tag)
		if err != nil {
			continue
		}

		switch field.Type.Kind() {
		case reflect.Int:
			state.SetGlobal(name, lua.LNumber(value.(int)))
		case reflect.String:
			state.SetGlobal(name, lua.LString(value.(string)))
		case reflect.Struct:
			table, err := Marshal(value)
			if err != nil {
				return err
			}
			state.SetGlobal(name, table)
		case reflect.Func:
			fn := func(s *lua.LState) int {
				props := []reflect.Value{}
				props = append(props, reflect.ValueOf(s.Get(1)))

				results := fromReal.Call(props)

				s.Push(results[0].Addr().Interface().(lua.LValue))
			
				return 1
			}
			state.SetGlobal(name, state.NewFunction(fn))
		default:
			return fmt.Errorf("unsupported type found: %s", field.Type.Name())
		}
	}
	return nil
}

func Eject(state *lua.LState, dest interface{}) error {
	destType := reflect.TypeOf(dest).Elem()
	destReal := reflect.ValueOf(dest).Elem()
	if destType.Kind() != reflect.Struct {
		return fmt.Errorf("unsupported value supplied")
	}

	for i := range destType.NumField() {
		field := destType.Field(i)
		value := destReal.Field(i)

		name, err := extarctLuaTagValue(field.Tag)
		if err != nil {
			continue
		}
		luaValue := state.GetGlobal(name)

		switch field.Type.Kind() {
		case reflect.Int:
			value.SetInt(int64(luaValue.(lua.LNumber)))
		case reflect.String:
			value.SetString(string(luaValue.(lua.LString)))
		case reflect.Struct:
			if err := Unmarshal(luaValue.(*lua.LTable), value.Addr().Interface()); err != nil {
				return err
			}
		case reflect.Func:
			luafn := luaValue.(*lua.LFunction)

			fn := func(arg interface{}) interface{} {
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
				return values[0]
			}
			value.Set(reflect.ValueOf(fn))
		default:
			return fmt.Errorf("unsupported type found: %s", field.Type.Name())
		}
	}
	return nil
}
