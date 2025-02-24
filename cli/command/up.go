package command

import (
	"context"
	"flag"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

func NewUpCommand() subcommands.Command {
	return &UpCommand{
		Container: infra.Default,
		handler: handle.New(infra.Default),
	}
}

type UpCommand struct {
	infra.Container
	handler handle.Handler
}

func (c *UpCommand) Name() string {
	return "up"
}

func (c *UpCommand) Synopsis() string {
	return "Up"
}

func (c *UpCommand) Usage() string {
	return "cywagon up [confpath...]\n"
}

func (c *UpCommand) SetFlags(_ *flag.FlagSet) {}

func (c *UpCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	paths := f.Args()

	if err := c.handler.Up(paths); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
