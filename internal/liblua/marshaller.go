package liblua

import (
	"fmt"

	"reflect"

	"github.com/fatih/structtag"
	lua "github.com/yuin/gopher-lua"
)

func extarctLuaTagValue(m reflect.StructTag) (string, error) {
	tags, err := structtag.Parse(string(m))
	if err != nil {
		return "", fmt.Errorf("lua tag not found: %s", err.Error())
	}
	value, err := tags.Get("lua")
	if err != nil {
		return "", fmt.Errorf("lua tag not found: %s", err.Error())
	}
	return value.Name, nil
}

func Marshal(from interface{}) (lua.LValue, error) {
	state := lua.NewState()
	table := state.NewTable()

	fromType := reflect.TypeOf(from)
	fromReal := reflect.ValueOf(from)

	if fromType.Kind() == reflect.Func {
		fn := func(s *lua.LState) int {
			table := s.ToTable(1)

			in0 := fromType.In(0)
			in0val := reflect.New(in0)

			dest := in0val.Interface()
			if err := Unmarshal(table, dest); err != nil {
				fmt.Println(err)
			}
			props := []reflect.Value{}
			props = append(props, in0val.Elem())
			fmt.Println("aaaaaa")

			results := fromReal.Call(props)

			luaVal, err := Marshal(results[0].Interface())
			if err != nil {
				fmt.Println(err)
				return 0
			}
			s.Push(luaVal)

			return 1
		}
		return state.NewFunction(fn), nil
	}

	if fromType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("unsupported value supplied %s", fromType)
	}

	for i := range fromType.NumField() {
		field := fromType.Field(i)
		value := fromReal.Field(i).Interface()

		name, err := extarctLuaTagValue(field.Tag)
		if err != nil {
			return nil, err
		}

		switch field.Type.Kind() {
		case reflect.Int:
			state.SetField(table, name, lua.LNumber(value.(int)))
		case reflect.String:
			state.SetField(table, name, lua.LString(value.(string)))
		default:
			return nil, fmt.Errorf("unsupported type found: %s", field.Type.Name())
		}
	}
	return table, nil
}

func Unmarshal(table lua.LValue, dest interface{}) error {
	state := lua.NewState()

	destType := reflect.TypeOf(dest).Elem()
	destReal := reflect.ValueOf(dest).Elem()

	for i := range destType.NumField() {
		field := destType.Field(i)
		value := destReal.Field(i)

		name, err := extarctLuaTagValue(field.Tag)
		if err != nil {
			return err
		}
		luaValue := state.GetField(table, name)

		switch field.Type.Kind() {
		case reflect.Int:
			value.SetInt(int64(luaValue.(lua.LNumber)))
		case reflect.String:
			value.SetString(string(luaValue.(lua.LString)))
		default:
			return fmt.Errorf("unsupported type found: %s", field.Type.Name())
		}
	}
	return nil
}
