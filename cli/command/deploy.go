package command

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

func NewDeployCommand() subcommands.Command {
	return &DeployCommand{
		Container: infra.Default,
		handler: handle.New(),
	}
}

type DeployCommand struct {
	infra.Container
	handler handle.Handler
}

func (c *DeployCommand) Name() string {
	return "deploy"
}

func (c *DeployCommand) Synopsis() string {
	return "Deploy"
}

func (c *DeployCommand) Usage() string {
	return "deploy <sitename> <path>\n"
}

func (c *DeployCommand) SetFlags(_ *flag.FlagSet) {}

func (c *DeployCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := c.validate(f.Args()); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	sitename := f.Arg(0)
	path := f.Arg(1)

	if err := c.handler.Deploy(sitename, path); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func (c *DeployCommand) validate(args []string) error {
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
