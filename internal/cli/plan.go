package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/internal/liblua"
	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/google/subcommands"
	lua "github.com/yuin/gopher-lua"
)

func newPlanCmd() *planCmd {
	return &planCmd{}
}

type planCmd struct {}

func (c *planCmd) Name() string {
	return "plan"
}

func (c *planCmd) Synopsis() string {
	return "plan"
}

func (c *planCmd) Usage() string {
	return "cywagon plan\n"
}

func (c *planCmd) SetFlags(f *flag.FlagSet) {}

func (c *planCmd) Execute(ctx context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	L := lua.NewState()
	defer L.Close()

	repos := repository.Use(ctx)
	scriptbytes, err := repos.Fs.Read("testdata/sites-enabled/example.lua")
	if err != nil {
		panic(err)
	}

	runner := liblua.NewRunner(string(scriptbytes))
	if err := runner.Run(); err != nil {
		panic(err)
	}

	fmt.Printf("hostname: %s\n", runner.GetString("hostname"))
	fmt.Printf("port: %d\n", runner.GetInt("port"))

	res := runner.S().NewTable()
	runner.S().SetField(res, "status", lua.LNumber(404))

	nextfn := L.NewFunction(Next)
	result, err := runner.RunFunction("handle", []lua.LValue{nextfn, nil, res})
	if err != nil {
		panic(err)
	}
	status := runner.S().GetField(result[0], "status")
	fmt.Printf("res: %+v\n", status)

	return subcommands.ExitSuccess
}

func Next(L *lua.LState) int {
	fmt.Println("this is next function")
	return 0
}
