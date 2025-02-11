package command

import (
	"context"
	"errors"
	"flag"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/google/subcommands"
)

var ErrPlanMissingRequiredFlagConf = errors.New("missing required flag: -conf")

func NewPlanCommand(container infra.Container) subcommands.Command {
	return &PlanCommand{
		Container: container,
	}
}

type PlanCommand struct {
	infra.Container
	conf string
}

func (c *PlanCommand) Name() string {
	return "plan"
}

func (c *PlanCommand) Synopsis() string {
	return "plan"
}

func (c *PlanCommand) Usage() string {
	return "cywagon plan\n"
}

func (c *PlanCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.conf, "conf", "", "conf files dir. required")
}

func (c *PlanCommand) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.conf == "" {
		c.Log.Error(ErrPlanMissingRequiredFlagConf)
		return subcommands.ExitFailure
	}

	if err := handle.Plan(c.Container, c.conf); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
