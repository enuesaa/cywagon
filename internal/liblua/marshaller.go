package liblua

import (
	"fmt"
	"strings"

	"reflect"

	"github.com/fatih/structtag"
	lua "github.com/yuin/gopher-lua"
)

func (r *Runner) extarctLuaTagValue(m reflect.StructTag, fieldName string) string {
	defaultName := strings.ToLower(fieldName)

	tags, err := structtag.Parse(string(m))
	if err != nil {
		return defaultName
	}
	value, err := tags.Get("lua")
	if err != nil {
		return defaultName
	}
	return value.Name
}

func (r *Runner) Marshal(from interface{}) (lua.LValue, error) {
	state := lua.NewState()
	table := state.NewTable()

	fromType := reflect.TypeOf(from)
	fromReal := reflect.ValueOf(from)
	kind := fromType.Kind()

	if kind == reflect.Func {
		arg := fromType.In(0)
		argReal := reflect.New(arg)

		fn := func(s *lua.LState) int {
			table := s.ToTable(1)

			if err := r.Unmarshal(table, argReal.Interface()); err != nil {
				fmt.Println(err)
			}
			results := fromReal.Call([]reflect.Value{
				argReal.Elem(),
			})
			result := results[0]

			luaResult, err := r.Marshal(result.Interface())
			if err != nil {
				fmt.Println(err)
				return 0
			}
			s.Push(luaResult)

			return 1
		}
		return state.NewFunction(fn), nil
	}

	if kind != reflect.Struct {
		return nil, fmt.Errorf("unsupported value supplied %s", fromType)
	}

	for i := range fromType.NumField() {
		field := fromType.Field(i)
		value := fromReal.Field(i).Interface()

		name := r.extarctLuaTagValue(field.Tag, field.Name)

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

func (r *Runner) Unmarshal(table lua.LValue, dest interface{}) error {
	state := lua.NewState()

	destType := reflect.TypeOf(dest).Elem()
	destReal := reflect.ValueOf(dest).Elem()

	for i := range destType.NumField() {
		field := destType.Field(i)
		value := destReal.Field(i)

		name := r.extarctLuaTagValue(field.Tag, field.Name)
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
