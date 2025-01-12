package cli

import (
	"context"
	"flag"
	"fmt"

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

	if err := L.DoString(string(scriptbytes)); err != nil {
		panic(err)
	}

	port := L.GetGlobal("port").(lua.LNumber)
	fmt.Printf("port: %d\n", port)

	hostname := L.GetGlobal("hostname").(lua.LString)
	fmt.Printf("hostname: %s\n", hostname)

	fn := L.GetGlobal("handle").(*lua.LFunction)

	res := L.NewTable()
	L.SetField(res, "status", lua.LNumber(404))

	nextfn := L.NewFunction(Next)

	_, err, values := L.Resume(lua.NewState(), fn, nextfn, nil, res)
	if err != nil {
		panic(err)
	}
	status := L.GetField(values[0], "status")
	fmt.Printf("res: %+v\n", status)

	return subcommands.ExitSuccess
}

func Next(L *lua.LState) int {
	fmt.Println("this is next function")
	return 0
}
