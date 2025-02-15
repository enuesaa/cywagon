package cli

import (
	"context"
	"flag"

	"github.com/enuesaa/cywagon/cli/command"
	"github.com/google/subcommands"
)

func Run() int {
	// cli
	subcommands.Register(command.NewPlanCommand(), "")
	subcommands.Register(command.NewStartCommand(), "")

	// parse
	flag.Parse()

	// execute
	status := subcommands.Execute(context.Background())

	return int(status)
}
