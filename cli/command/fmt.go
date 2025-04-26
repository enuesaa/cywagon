package command

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

func NewFmtCommand() subcommands.Command {
	return &FmtCommand{
		Container: infra.Default,
		handler:   handle.New(),
	}
}

type FmtCommand struct {
	infra.Container
	handler handle.Handler
}

func (c *FmtCommand) Name() string {
	return "fmt"
}

func (c *FmtCommand) Synopsis() string {
	return "[Experimental] Fmt"
}

func (c *FmtCommand) Usage() string {
	return "fmt <workdir>\n"
}

func (c *FmtCommand) SetFlags(_ *flag.FlagSet) {}

func (c *FmtCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	if err := c.ValidateArgs(f.Args()); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	path := f.Arg(0)

	if err := c.handler.Format(path); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func (c *FmtCommand) ValidateArgs(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("required arguments missing: path")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments found")
	}
	return nil
}