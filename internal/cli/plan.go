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

type PlanCmd struct {
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

func (c *PlanCmd) Execute(ctx context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	repos := repository.Use(ctx)

	if c.conf == "" {
		repos.Log.Error(ErrPlanMissingRequiredFlagConf)
		return subcommands.ExitFailure
	}

	if err := usecase.Plan(ctx, c.conf); err != nil {
		repos.Log.Error(err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
