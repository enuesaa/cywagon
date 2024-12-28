package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/internal/engine"
	"github.com/google/subcommands"
)

type createCmd struct {}

func (c *createCmd) Name() string {
	return "create"
}

func (c *createCmd) Synopsis() string {
	return "create"
}

func (c *createCmd) Usage() string {
	return `create:
	Create
  `
}

func (c *createCmd) SetFlags(f *flag.FlagSet) {}

func (c *createCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := engine.SendCreate(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
