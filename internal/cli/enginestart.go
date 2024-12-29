package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/internal/eng"
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

func (c *engineStartCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := eng.RunEngine(ctx); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return subcommands.ExitFailure
	}
	
	return subcommands.ExitSuccess
}
