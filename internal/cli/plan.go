package cli

import (
	"context"
	"errors"
	"flag"

	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/enuesaa/cywagon/internal/usecase"
	"github.com/google/subcommands"
)

var ErrPlanMissingRequiredFlagConf = errors.New("missing required flag: -conf")

func NewPlanCmd(repos repository.Repos) subcommands.Command {
	return &PlanCmd{
		repos: repos,
	}
}

type PlanCmd struct {
	repos repository.Repos
	conf string
}

func (c *PlanCmd) Name() string {
	return "plan"
}

func (c *PlanCmd) Synopsis() string {
	return "plan"
}

func (c *PlanCmd) Usage() string {
	return "cywagon plan\n"
}

func (c *PlanCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.conf, "conf", "", "conf files dir. required")
}

func (c *PlanCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.conf == "" {
		c.repos.Log.Error(ErrPlanMissingRequiredFlagConf)
		return subcommands.ExitFailure
	}

	if err := usecase.Plan(c.repos, c.conf); err != nil {
		c.repos.Log.Error(err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
