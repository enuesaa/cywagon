package cli

import (
	"context"
	"flag"
	"log"

	"github.com/enuesaa/cywagon/internal/engine"
	"github.com/google/subcommands"
)

type engineStartCmd struct {}

func (c *engineStartCmd) Name() string {
	return "engine-start"
}

func (c *engineStartCmd) Synopsis() string {
	return "Start engine"
}

func (c *engineStartCmd) Usage() string {
	return `engine-start:
	Start engine.
  `
}

func (c *engineStartCmd) SetFlags(f *flag.FlagSet) {}

func (c *engineStartCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := engine.RunEngine(); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	
	return subcommands.ExitSuccess
}
