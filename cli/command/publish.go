package command

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

func NewPublishCommand() subcommands.Command {
	return &PublishCommand{
		Container: infra.Default,
		handler: handle.New(),
	}
}

type PublishCommand struct {
	infra.Container
	handler handle.Handler
}

func (c *PublishCommand) Name() string {
	return "publish"
}

func (c *PublishCommand) Synopsis() string {
	return "Publish"
}

func (c *PublishCommand) Usage() string {
	return "publish <sitename> <deploymentId>\n"
}

func (c *PublishCommand) SetFlags(_ *flag.FlagSet) {}

func (c *PublishCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := c.validate(f.Args()); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	sitename := f.Arg(0)
	deploymentId := f.Arg(1)

	if err := c.handler.Publish(sitename, deploymentId); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func (c *PublishCommand) validate(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("required arguments missing: sitename, deploymentId")
	}
	if len(args) == 1 {
		return fmt.Errorf("required arguments missing: deploymentId")
	}
	if len(args) > 2 {
		return fmt.Errorf("too many arguments found")
	}
	return nil
}
