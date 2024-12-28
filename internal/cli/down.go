package cli

import (
	"context"
	"flag"
	"log"

	"github.com/enuesaa/cywagon/internal/engine"
	"github.com/google/subcommands"
)

type downCmd struct {}

func (c *downCmd) Name() string {
	return "down"
}

func (c *downCmd) Synopsis() string {
	return "down"
}

func (c *downCmd) Usage() string {
	return `down:
	Down engine.
  `
}

func (c *downCmd) SetFlags(f *flag.FlagSet) {}

func (c *downCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := engine.Down(); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	return subcommands.ExitSuccess
}
