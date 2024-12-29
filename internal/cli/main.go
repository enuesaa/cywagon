package cli

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

func Run(ctx context.Context) int {
	// cli
	subcommands.Register(&upCmd{}, "")
	subcommands.Register(&createCmd{}, "")
	subcommands.Register(&engineStartCmd{}, "")
	subcommands.Register(&downCmd{}, "")

	// parse
	flag.Parse()

	// execute
	status := subcommands.Execute(ctx)

	return int(status)
}
