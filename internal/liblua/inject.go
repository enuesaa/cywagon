package liblua

import (
	"fmt"
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

type Fn func(res interface{}, args ...interface{}) error

func (r *Runner) Inject(from interface{}) error {
	fromType := reflect.TypeOf(from)
	fromReal := reflect.ValueOf(from)
	if fromType.Kind() != reflect.Struct {
		return fmt.Errorf("unsupported value supplied")
	}

	for i := range fromType.NumField() {
		field := fromType.Field(i)
		value := fromReal.Field(i).Interface()

		name := r.extarctVarName(field.Tag, field.Name)

		switch field.Type.Kind() {
		case reflect.Int:
			r.state.SetGlobal(name, lua.LNumber(value.(int)))
		case reflect.String:
			r.state.SetGlobal(name, lua.LString(value.(string)))
		case reflect.Bool:
			r.state.SetGlobal(name, lua.LBool(value.(bool)))
		case reflect.Struct:
			table, err := r.Marshal(value)
			if err != nil {
				return err
			}
			r.state.SetGlobal(name, table)
		case reflect.Func:
			// pass
		default:
			return fmt.Errorf("unsupported type found: %s", field.Type.Name())
		}
	}
	return nil
}

func (r *Runner) Eject(dest interface{}) error {
	destType := reflect.TypeOf(dest).Elem()
	destReal := reflect.ValueOf(dest).Elem()
	if destType.Kind() != reflect.Struct {
		return fmt.Errorf("unsupported value supplied")
	}

	for i := range destType.NumField() {
		field := destType.Field(i)
		value := destReal.Field(i)

		name := r.extarctVarName(field.Tag, field.Name)
		luaValue := r.state.GetGlobal(name)

		switch field.Type.Kind() {
		case reflect.Int:
			value.SetInt(int64(luaValue.(lua.LNumber)))
		case reflect.String:
			value.SetString(string(luaValue.(lua.LString)))
		case reflect.Bool:
			value.SetBool(bool(luaValue.(lua.LBool)))
		case reflect.Struct:
			if err := r.Unmarshal(luaValue.(*lua.LTable), value.Addr().Interface()); err != nil {
				return err
			}
		case reflect.Func:
			luafn := luaValue.(*lua.LFunction)

			fn := func(res interface{}, args ...interface{}) error {
				luaArgs := []lua.LValue{}
				for _, arg := range args {
					luaArg, err := r.Marshal(arg)
					if err != nil {
						return err
					}
					luaArgs = append(luaArgs, luaArg)
				}

				_, err, values := r.state.Resume(lua.NewState(), luafn, luaArgs...)
				if err != nil {
					return err
				}
				if len(values) == 0 {
					return nil
				}
				return r.Unmarshal(values[0], res)
			}
			value.Set(reflect.ValueOf(fn))
		default:
			return fmt.Errorf("unsupported type found: %s", field.Type.Name())
		}
	}
	return nil
}
