package cli

import (
	"context"
	"flag"
	"log"

	"github.com/enuesaa/cywagon/internal/engine"
	"github.com/google/subcommands"
)

type upCmd struct {}

func (c *upCmd) Name() string {
	return "up"
}

func (c *upCmd) Synopsis() string {
	return "up"
}

func (c *upCmd) Usage() string {
	return `up:
	Start engine.
  `
}

func (c *upCmd) SetFlags(f *flag.FlagSet) {}

func (c *upCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := engine.Up(); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	
	return subcommands.ExitSuccess
}
