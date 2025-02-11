package cli

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

func Run(ctx context.Context) int {
	// cli
	subcommands.Register(&PlanCmd{}, "")
	subcommands.Register(&StartCmd{}, "")

	// parse
	flag.Parse()

	// execute
	status := subcommands.Execute(ctx)

	return int(status)
}
