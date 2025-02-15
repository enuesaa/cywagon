package command

import (
	"context"
	"errors"
	"flag"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

var ErrUpMissingRequiredFlagConf = errors.New("missing required flag: -conf")

func NewUpCommand() subcommands.Command {
	return &UpCommand{
		Container: infra.Default,
	}
}

type UpCommand struct {
	infra.Container
	conf string
}

func (c *UpCommand) Name() string {
	return "up"
}

func (c *UpCommand) Synopsis() string {
	return "Up"
}

func (c *UpCommand) Usage() string {
	return "cywagon up\n"
}

func (c *UpCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.conf, "conf", "", "conf files dir. required")
}

func (c *UpCommand) Execute(ctx context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := c.validate(); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	if err := handle.Up(c.conf); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func (c *UpCommand) validate() error {
	if c.conf == "" {
		return ErrUpMissingRequiredFlagConf
	}
	return nil
}
