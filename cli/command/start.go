package cli

import (
	"context"
	"errors"
	"flag"

	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/enuesaa/cywagon/internal/usecase"
	"github.com/google/subcommands"
)

var ErrStartMissingRequiredFlagConf = errors.New("missing required flag: -conf")

func NewStartCmd(repos repository.Repos) subcommands.Command {
	return &StartCmd{
		repos: repos,
	}
}

type StartCmd struct {
	repos repository.Repos
	conf string
}

func (c *StartCmd) Name() string {
	return "start"
}

func (c *StartCmd) Synopsis() string {
	return "Start"
}

func (c *StartCmd) Usage() string {
	return "cywagon start\n"
}

func (c *StartCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.conf, "conf", "", "conf files dir. required")
}

func (c *StartCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.conf == "" {
		c.repos.Log.Error(ErrStartMissingRequiredFlagConf)
		return subcommands.ExitFailure
	}

	if err := usecase.Start(c.repos, c.conf); err != nil {
		c.repos.Log.Error(err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
