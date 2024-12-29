package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/internal/msg"
	"github.com/google/subcommands"
)

type createCmd struct {}

func (c *createCmd) Name() string {
	return "create"
}

func (c *createCmd) Synopsis() string {
	return "Create"
}

func (c *createCmd) Usage() string {
	return "cywagon create\n"
}

func (c *createCmd) SetFlags(f *flag.FlagSet) {}

func (c *createCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	sender := msg.Sender{}
	if err := sender.SendCreateMessage("aaa"); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
