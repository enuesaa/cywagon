package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/cli/command"
	"github.com/google/subcommands"
)

var versionFlag = flag.Bool("version", false, "Print version")

func Run() int {
	// cli
	subcommands.Register(command.NewCheckCommand(), "")
	subcommands.Register(command.NewUpCommand(), "")

	// parse
	flag.Parse()

	if *versionFlag {
		fmt.Println("0.0.1")
		return 0
	}

	// execute
	status := subcommands.Execute(context.Background())

	return int(status)
}
