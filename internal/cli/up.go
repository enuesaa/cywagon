package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/internal/eng"
	"github.com/enuesaa/cywagon/internal/engctl"
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
	if c.foreground {
		if err := eng.Up(ctx); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return subcommands.ExitFailure
		}
		return subcommands.ExitSuccess
	}

	if err := engctl.Up(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
