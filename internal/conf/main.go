package conf

import (
	"context"
	"fmt"

	"github.com/enuesaa/cywagon/internal/liblua"
	"github.com/enuesaa/cywagon/internal/repository"
	lua "github.com/yuin/gopher-lua"
)

func Parse(ctx context.Context) error {
	repos := repository.Use(ctx)
	scriptbytes, err := repos.Fs.Read("testdata/sites-enabled/example.lua")
	if err != nil {
		return err
	}

	runner := liblua.NewRunner(string(scriptbytes))
	if err := runner.Run(); err != nil {
		return err
	}

	fmt.Printf("hostname: %s\n", runner.GetString("hostname"))
	fmt.Printf("port: %d\n", runner.GetInt("port"))

	res := runner.S().NewTable()
	runner.S().SetField(res, "status", lua.LNumber(404))

	nextfn := runner.S().NewFunction(Next)
	result, err := runner.RunFunction("handle", []lua.LValue{nextfn, nil, res})
	if err != nil {
		return err
	}
	status := runner.S().GetField(result[0], "status")
	fmt.Printf("res: %+v\n", status)

	return nil
}

type Config struct {}

func Next(L *lua.LState) int {
	fmt.Println("this is next function")
	return 0
}
