package cli

import (
	"context"
	"flag"

	"github.com/enuesaa/cywagon/cli/command"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/google/subcommands"
)

func Run(container infra.Container) int {
	// cli
	subcommands.Register(command.NewPlanCmd(container), "")
	subcommands.Register(command.NewStartCmd(container), "")

	// parse
	flag.Parse()

	// execute
	status := subcommands.Execute(context.Background())

	return int(status)
}
