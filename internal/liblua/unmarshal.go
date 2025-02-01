package liblua

import (
	"fmt"

	"reflect"

	"github.com/fatih/structtag"
	lua "github.com/yuin/gopher-lua"
)

func Unmarshal(table *lua.LTable, dest interface{}) error {
	state := lua.NewState()

	target := reflect.TypeOf(dest).Elem()
	targetValue := reflect.ValueOf(dest).Elem()

	for i := range target.NumField() {
		field := target.Field(i)

		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			return err
		}
		luaTag, err := tags.Get("lua")
		if err != nil {
			return err
		}

		fieldType := field.Type.Name()
		refValue := targetValue.FieldByName(field.Name)

		switch fieldType {
		case "int":
			val := state.GetField(table, luaTag.Name).(lua.LNumber)
			refValue.SetInt(int64(val))
		case "string":
			val := state.GetField(table, luaTag.Name).(lua.LString)
			refValue.SetString(string(val))
		default:
			return fmt.Errorf("unknown")
		}
	}
	return nil
}
