package command

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

func NewCheckCommand() subcommands.Command {
	return &CheckCommand{
		Container: infra.Default,
		handler:   handle.New(),
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
	return "check <confpath>\n"
}

func (c *CheckCommand) SetFlags(_ *flag.FlagSet) {}

func (c *CheckCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	if err := c.ValidateArgs(f.Args()); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	path := f.Arg(0)

	if err := c.handler.Check(path); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func (c *CheckCommand) ValidateArgs(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("required arguments missing: path")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments found")
	}
	return nil
}

