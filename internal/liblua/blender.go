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
		default:
			return fmt.Errorf("unsupported type found: %s", field.Type.Name())
		}
	}
	return nil
}
