package command

import (
	"context"
	"errors"
	"flag"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

var ErrPlanMissingRequiredFlagConf = errors.New("missing required flag: -conf")

func NewPlanCommand() subcommands.Command {
	return &PlanCommand{
		Container: infra.Default,
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
	if err := c.validate(); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	if err := handle.Plan(c.conf); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func (c *PlanCommand) validate() error {
	if c.conf == "" {
		return ErrPlanMissingRequiredFlagConf
	}
	return nil
}
