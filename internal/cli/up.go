package cli

import (
	"context"
	"flag"

	"github.com/enuesaa/cywagon/internal/eng"
	"github.com/enuesaa/cywagon/internal/engctl"
	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/google/subcommands"
)

func newUpCmd() *upCmd {
	return &upCmd{
		foreground: false,
	}
}

type upCmd struct {
	foreground bool
}

func (c *upCmd) Name() string {
	return "up"
}

func (c *upCmd) Synopsis() string {
	return "Up"
}

func (c *upCmd) Usage() string {
	return "cywagon up\n"
}

func (c *upCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.foreground, "foreground", false, "run foreground")
}

func (c *upCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	repos := repository.Use(ctx)

	if c.foreground {
		if err := eng.Up(ctx); err != nil {
			repos.Log.PrintErr(err)
			return subcommands.ExitFailure
		}
	} else {
		if err := engctl.Up(ctx); err != nil {
			repos.Log.PrintErr(err)
			return subcommands.ExitFailure
		}
	}
	return subcommands.ExitSuccess
}
