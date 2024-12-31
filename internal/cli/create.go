package cli

import (
	"context"
	"flag"

	"github.com/enuesaa/cywagon/internal/engctl"
	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/google/subcommands"
)

func newCreateCmd() *createCmd {
	return &createCmd{}
}

type createCmd struct {}

func (c *createCmd) Name() string {
	return "create"
}

func (c *createCmd) Synopsis() string {
	return "Create"
}

func (c *createCmd) Usage() string {
	return "cywagon create\n"
}

func (c *createCmd) SetFlags(f *flag.FlagSet) {}

func (c *createCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	repos := repository.Use(ctx)

	if err := engctl.SendCreateMessage(ctx); err != nil {
		repos.Log.PrintErr(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
