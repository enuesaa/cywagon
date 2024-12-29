package cli

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

func Run(ctx context.Context) int {
	// cli
	subcommands.Register(newUpCmd(), "")
	subcommands.Register(newDownCmd(), "")
	subcommands.Register(newCreateCmd(), "")

	// parse
	flag.Parse()

	// execute
	status := subcommands.Execute(ctx)

	return int(status)
}
