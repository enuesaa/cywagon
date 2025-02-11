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

func NewPlanCommand(ctn infra.Container) subcommands.Command {
	return &PlanCommand{
		ctn: ctn,
	}
}

type PlanCommand struct {
	ctn infra.Container
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
		c.ctn.Log.Error(ErrPlanMissingRequiredFlagConf)
		return subcommands.ExitFailure
	}

	if err := handle.Plan(c.ctn, c.conf); err != nil {
		c.ctn.Log.Error(err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
