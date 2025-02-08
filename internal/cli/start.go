package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/enuesaa/cywagon/internal/usecase"
	"github.com/google/subcommands"
)

func newStartCmd() *startCmd {
	return &startCmd{
		conf: ".",
	}
}

type startCmd struct {
	conf string
}

func (c *startCmd) Name() string {
	return "start"
}

func (c *startCmd) Synopsis() string {
	return "Start"
}

func (c *startCmd) Usage() string {
	return "cywagon start\n"
}

func (c *startCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.conf, "conf", "", "conf files dir. required")
}

func (c *startCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	repos := repository.Use(ctx)

	if c.conf == "" {
		err := fmt.Errorf("missing required flag: -conf")
		repos.Log.Error(err)
		return subcommands.ExitFailure
	}

	if err := usecase.Start(ctx, c.conf); err != nil {
		repos.Log.Error(err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
