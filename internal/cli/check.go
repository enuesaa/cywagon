package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/google/subcommands"
	lua "github.com/yuin/gopher-lua"
)

func newCheckCmd() *checkCmd {
	return &checkCmd{}
}

type checkCmd struct {}

func (c *checkCmd) Name() string {
	return "check"
}

func (c *checkCmd) Synopsis() string {
	return "check"
}

func (c *checkCmd) Usage() string {
	return "cywagon check\n"
}

func (c *checkCmd) SetFlags(f *flag.FlagSet) {}

func (c *checkCmd) Execute(ctx context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
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

	co, cocancel := L.NewThread()
	defer cocancel()
	fn := L.GetGlobal("handle").(*lua.LFunction)
	_, err, values := L.Resume(co, fn)
	if err != nil {
		panic(err)
	}
	fmt.Printf("res: %v\n", values)
	fmt.Printf("res\n")

	// return subcommands.ExitSuccess
	return 0
}
