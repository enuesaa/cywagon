package command

import (
	"context"
	"errors"
	"flag"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/google/subcommands"
)

var ErrStartMissingRequiredFlagConf = errors.New("missing required flag: -conf")

func NewStartCommand(container infra.Container) subcommands.Command {
	return &StartCommand{
		Container: container,
	}
}

type StartCommand struct {
	infra.Container
	conf string
}

func (c *StartCommand) Name() string {
	return "start"
}

func (c *StartCommand) Synopsis() string {
	return "Start"
}

func (c *StartCommand) Usage() string {
	return "cywagon start\n"
}

func (c *StartCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.conf, "conf", "", "conf files dir. required")
}

func (c *StartCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.conf == "" {
		c.Log.Error(ErrStartMissingRequiredFlagConf)
		return subcommands.ExitFailure
	}

	if err := handle.Start(c.Container, c.conf); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
