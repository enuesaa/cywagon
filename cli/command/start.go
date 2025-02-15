package command

import (
	"context"
	"errors"
	"flag"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

var ErrStartMissingRequiredFlagConf = errors.New("missing required flag: -conf")

func NewStartCommand() subcommands.Command {
	return &StartCommand{
		Container: infra.Default,
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

func (c *StartCommand) Execute(ctx context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := c.validate(); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	if err := handle.Start(c.conf); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func (c *StartCommand) validate() error {
	if c.conf == "" {
		return ErrStartMissingRequiredFlagConf
	}
	return nil
}
