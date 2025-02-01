package cli

import (
	"context"
	"flag"
	"log"

	"github.com/enuesaa/cywagon/internal/usecase"
	"github.com/google/subcommands"
)

func newPlanCmd() *planCmd {
	return &planCmd{
		conf: ".",
	}
}

type planCmd struct {
	conf string
}

func (c *planCmd) Name() string {
	return "plan"
}

func (c *planCmd) Synopsis() string {
	return "plan"
}

func (c *planCmd) Usage() string {
	return "cywagon plan\n"
}

func (c *planCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.conf, ".", ".", "conf files dir")
}

func (c *planCmd) Execute(ctx context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := usecase.Plan(ctx, c.conf); err != nil {
		log.Printf("Error: %s", err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
