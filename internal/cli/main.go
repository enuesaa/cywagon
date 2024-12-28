package cli

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

func Run() {
	subcommands.Register(&upCmd{}, "")
	subcommands.Register(&createCmd{}, "")
	subcommands.Register(&engineStartCmd{}, "")
	subcommands.Register(&downCmd{}, "")
  
	flag.Parse()

	ctx := context.Background()
	status := subcommands.Execute(ctx)

	os.Exit(int(status))
}
