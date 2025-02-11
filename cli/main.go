package cli

import (
	"context"
	"flag"

	"github.com/enuesaa/cywagon/cli/command"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

func Run(ctn infra.Container) int {
	// cli
	subcommands.Register(command.NewPlanCommand(ctn), "")
	subcommands.Register(command.NewStartCommand(ctn), "")

	// parse
	flag.Parse()

	// execute
	status := subcommands.Execute(context.Background())

	return int(status)
}
