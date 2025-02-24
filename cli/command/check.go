package command

import (
	"context"
	"flag"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

func NewCheckCommand(container infra.Container) subcommands.Command {
	return &CheckCommand{
		Container: container,
		handler: handle.New(container),
	}
}

type CheckCommand struct {
	infra.Container
	handler handle.Handler
}

func (c *CheckCommand) Name() string {
	return "check"
}

func (c *CheckCommand) Synopsis() string {
	return "Check"
}

func (c *CheckCommand) Usage() string {
	return "cywagon check [confpath...]\n"
}

func (c *CheckCommand) SetFlags(_ *flag.FlagSet) {}

func (c *CheckCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	paths := f.Args()

	if err := c.handler.Check(paths); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
