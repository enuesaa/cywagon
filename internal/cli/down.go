package cli

import (
	"context"
	"flag"

	"github.com/enuesaa/cywagon/internal/engctl"
	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/google/subcommands"
)

func newDownCmd() *downCmd {
	return &downCmd{}
}

type downCmd struct {}

func (c *downCmd) Name() string {
	return "down"
}

func (c *downCmd) Synopsis() string {
	return "Down"
}

func (c *downCmd) Usage() string {
	return "cywagon down\n"
}

func (c *downCmd) SetFlags(f *flag.FlagSet) {}

func (c *downCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	repos := repository.Use(ctx)

	if err := engctl.Down(ctx); err != nil {
		repos.Log.PrintErr(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
