package cli

import (
	"context"
	"flag"

	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/google/subcommands"
)

func Run(repos repository.Repos) int {
	// cli
	subcommands.Register(NewPlanCmd(repos), "")
	subcommands.Register(NewStartCmd(repos), "")

	// parse
	flag.Parse()

	// execute
	status := subcommands.Execute(context.Background())

	return int(status)
}
