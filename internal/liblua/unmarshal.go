package liblua

import (
	"fmt"

	"reflect"

	"github.com/fatih/structtag"
	lua "github.com/yuin/gopher-lua"
)

func Unmarshal(table *lua.LTable, dest interface{}) error {
	state := lua.NewState()

	target := reflect.TypeOf(dest)
	targetValue := reflect.ValueOf(dest)
	target = target.Elem()
	targetValue = targetValue.Elem()

	for i := range target.NumField() {
		field := target.Field(i)
		refvalue := targetValue.FieldByName(field.Name)

		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			return err
		}

		luaTag, err := tags.Get("lua")
		if err != nil {
			return err
		}

		if field.Type.Name() == "int" {
			val := state.GetField(table, luaTag.Name).(lua.LNumber)
			refvalue.SetInt(int64(val))
		} else if field.Type.Name() == "string" {
			val := state.GetField(table, luaTag.Name).(lua.LString)
			refvalue.SetString(string(val))
		} else {
			return fmt.Errorf("unknown")
		}
	}
	return nil
}
