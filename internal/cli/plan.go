package cli

import (
	"context"
	"flag"
	"log"

	"github.com/enuesaa/cywagon/internal/conf"
	"github.com/google/subcommands"
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
	if err := conf.Parse(ctx); err != nil {
		log.Printf("Error: %s", err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
