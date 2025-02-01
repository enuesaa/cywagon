package cli

import (
	"context"
	"flag"

	"github.com/enuesaa/cywagon/internal/conf"
	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/google/subcommands"
)

func newPlanCmd() *planCmd {
	return &planCmd{}
}

type planCmd struct{}

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
	repos := repository.Use(ctx)

	config, err := libconf.Parse(ctx, "testdata/sites-enabled/example.lua")
	if err != nil {
		repos.Log.PrintErr(err)
		return subcommands.ExitFailure
	}
	repos.Log.Print("hostname: %s\n", config.Hostname)
	repos.Log.Print("port: %d\n", config.Port)

	if err := config.RunHandler(); err != nil {
		repos.Log.PrintErr(err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
