package liblua

import (
	"fmt"

	"reflect"

	"github.com/fatih/structtag"
	lua "github.com/yuin/gopher-lua"
)

func Marshal(from interface{}) (*lua.LTable, error) {
	state := lua.NewState()
	ret := state.NewTable()

	target := reflect.TypeOf(from)
	targetValue := reflect.ValueOf(from)

	for i := range target.NumField() {
		field := target.Field(i)
		value := targetValue.Field(i).Interface()

		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			return nil, err
		}

		luaTag, err := tags.Get("lua")
		if err != nil {
			return nil, err
		}

		fieldType := field.Type.Name()
		switch fieldType {
		case "int":
			state.SetField(ret, luaTag.Name, lua.LNumber(value.(int)))
		case "string":
			state.SetField(ret, luaTag.Name, lua.LString(value.(string)))
		default:
			return nil, fmt.Errorf("unknown")
		}
	}
	return ret, nil
}
