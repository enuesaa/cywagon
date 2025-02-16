package command

import (
	"context"
	"errors"
	"flag"

	"github.com/enuesaa/cywagon/cli/handle"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

var ErrCheckMissingRequiredFlagConf = errors.New("missing required flag: -conf")

func NewCheckCommand() subcommands.Command {
	return &CheckCommand{
		Container: infra.Default,
	}
}

type CheckCommand struct {
	infra.Container
	conf string
}

func (c *CheckCommand) Name() string {
	return "check"
}

func (c *CheckCommand) Synopsis() string {
	return "check"
}

func (c *CheckCommand) Usage() string {
	return "cywagon check\n"
}

func (c *CheckCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.conf, "conf", "", "conf files dir. required")
}

func (c *CheckCommand) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	handler := handle.New()

	if err := c.validate(); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	if err := handler.Check(c.conf); err != nil {
		c.Log.Error(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func (c *CheckCommand) validate() error {
	if c.conf == "" {
		return ErrCheckMissingRequiredFlagConf
	}
	return nil
}
