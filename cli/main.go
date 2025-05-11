package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/enuesaa/cywagon/cli/command"
	"github.com/google/subcommands"
)

var versionFlag = flag.Bool("version", false, "Print version")
var helpFlag = flag.Bool("help", false, "Print command usage")

func Run() int {
	subcommands.DefaultCommander.Explain = FprintHelpText
	subcommands.Register(command.NewUpCommand(), "")
	subcommands.Register(command.NewCheckCommand(), "")
	subcommands.Register(command.NewFmtCommand(), "")

	// parse
	flag.Parse()

	if *versionFlag {
		fmt.Println("0.0.4")
		return 0
	}

	// execute
	status := subcommands.Execute(context.Background())

	if *helpFlag {
		status = subcommands.ExitSuccess
	}
	return int(status)
}
