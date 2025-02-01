package liblua

import (
	"log"

	"fmt"
	"reflect"

	"github.com/fatih/structtag"
	lua "github.com/yuin/gopher-lua"
)

func Marshal(from interface{}) *lua.LTable {
	state := lua.NewState()
	ret := state.NewTable()

	target := reflect.TypeOf(from)
	targetValue := reflect.ValueOf(from)

	for i := range target.NumField() {
		field := target.Field(i)
		value := targetValue.Field(i).Interface()

		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			continue
		}

		luaTag, err := tags.Get("lua")
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			continue
		}
		fmt.Printf("MARSHAL: %s ===>>> %+v\n", luaTag.Name, target.Field(i))

		if field.Type.Name() == "int" {
			state.SetField(ret, luaTag.Name, lua.LNumber(value.(int)))
		} else if field.Type.Name() == "string" {
			state.SetField(ret, luaTag.Name, lua.LString(value.(string)))
		} else {
			log.Printf("Error: unknown\n")
		}
	}

	return ret
}
