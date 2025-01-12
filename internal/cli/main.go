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
	subcommands.Register(newCheckCmd(), "")

	// parse
	flag.Parse()

	// execute
	status := subcommands.Execute(ctx)

	return int(status)
}
