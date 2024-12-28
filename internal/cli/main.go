package cli

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

func Run() int {
	subcommands.Register(&upCmd{}, "")
	subcommands.Register(&helloCmd{}, "")
	subcommands.Register(&engineStartCmd{}, "")
  
	flag.Parse()

	ctx := context.Background()
	status := subcommands.Execute(ctx)

	return int(status)
}
