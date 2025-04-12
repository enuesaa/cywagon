package command

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

func NewReloadCommand() subcommands.Command {
	return &ReloadCommand{
		Container: infra.Default,
		handler: handle.New(),
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
	return "reload <sitename> <path>\n"
}

func (c *ReloadCommand) SetFlags(_ *flag.FlagSet) {}

func (c *ReloadCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := c.validate(f.Args()); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	sitename := f.Arg(0)
	path := f.Arg(1)

	if err := c.handler.Reload(sitename, path); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func (c *ReloadCommand) validate(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("required arguments missing: sitename, path")
	}
	if len(args) == 1 {
		return fmt.Errorf("required arguments missing: path")
	}
	if len(args) > 2 {
		return fmt.Errorf("too many arguments found")
	}
	return nil
}
