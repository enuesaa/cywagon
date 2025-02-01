package cli

import (
	"context"
	"flag"
	"log"

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
	return "up"
}

func (c *startCmd) Synopsis() string {
	return "Up"
}

func (c *startCmd) Usage() string {
	return "cywagon up\n"
}

func (c *startCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.conf, ".", ".", "conf files dir")
}

func (c *startCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := usecase.Start(ctx, c.conf); err != nil {
		log.Printf("Error: %s", err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
