package cli

import (
	"flag"
	"os"

	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/google/subcommands"
)

func Run() {
	subcommands.Register(&upCmd{}, "")
	subcommands.Register(&createCmd{}, "")
	subcommands.Register(&engineStartCmd{}, "")
	subcommands.Register(&downCmd{}, "")
  
	flag.Parse()

	ctx := repository.NewContext()
	status := subcommands.Execute(ctx)

	os.Exit(int(status))
}
