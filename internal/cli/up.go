package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/internal/engctl"
	"github.com/google/subcommands"
)

type upCmd struct {}

func (c *upCmd) Name() string {
	return "up"
}

func (c *upCmd) Synopsis() string {
	return "Up"
}

func (c *upCmd) Usage() string {
	return "cywagon up\n"
}

func (c *upCmd) SetFlags(f *flag.FlagSet) {}

func (c *upCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := engctl.Up(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
