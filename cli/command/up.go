package command

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

func NewUpCommand() subcommands.Command {
	return &UpCommand{
		Container: infra.Default,
		handler:   handle.New(),
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
	return "cywagon up <workdir>\n"
}

func (c *UpCommand) SetFlags(_ *flag.FlagSet) {}

func (c *UpCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	if err := c.ValidateArgs(f.Args()); err != nil {
		c.Ps.PrintErr(err)
		return subcommands.ExitFailure
	}
	path := f.Arg(0)

	if err := c.handler.Up(path); err != nil {
		c.Ps.PrintErr(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func (c *UpCommand) ValidateArgs(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("required arguments missing: path")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments found")
	}
	return nil
}
