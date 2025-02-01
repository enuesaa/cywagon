package cli

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

func Run(ctx context.Context) int {
	// cli
	subcommands.Register(newPlanCmd(), "")
	subcommands.Register(newStartCmd(), "")

	// parse
	flag.Parse()

	// execute
	status := subcommands.Execute(ctx)

	return int(status)
}
