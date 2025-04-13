package command

import (
	"context"
	"flag"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

func NewReloadCommand() subcommands.Command {
	return &ReloadCommand{
		Container: infra.Default,
		handler:   handle.New(),
	}
}

type ReloadCommand struct {
	infra.Container
	handler handle.Handler
}

func (c *ReloadCommand) Name() string {
	return "reload"
}

func (c *ReloadCommand) Synopsis() string {
	return "Reload"
}

func (c *ReloadCommand) Usage() string {
	return "reload <confpath>\n"
}

func (c *ReloadCommand) SetFlags(_ *flag.FlagSet) {}

func (c *ReloadCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	if err := c.handler.ValidateArgs(f.Args()); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	path := f.Arg(0)

	if err := c.handler.Reload("", path); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
