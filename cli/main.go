package cli

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/enuesaa/cywagon/cli/command"
	"github.com/google/subcommands"
)

var versionFlag = flag.Bool("version", false, "Print version")
var helpFlag    = flag.Bool("help", false, "Print command usage")

func Run() int {
	// cli
	subcommands.DefaultCommander.Explain = Explain
	subcommands.Register(command.NewUpCommand(), "")
	subcommands.Register(command.NewCheckCommand(), "")
	subcommands.Register(command.NewReloadCommand(), "")

	// parse
	flag.Parse()

	if *versionFlag {
		fmt.Println("0.0.1")
		return 0
	}

	// execute
	status := subcommands.Execute(context.Background())

	if *helpFlag {
		status = subcommands.ExitSuccess
	}
	return int(status)
}

func Explain(w io.Writer) {
	cdr := subcommands.DefaultCommander

	fmt.Fprintf(w, "Usage: %s <subcommand>\n\n", cdr.Name())

	fmt.Fprintf(w, "Commands:\n")
	cdr.VisitCommands(func(_ *subcommands.CommandGroup, c subcommands.Command) {
		fmt.Fprintf(w, "\t%s      \t%s\n", c.Name(), c.Synopsis())
	})

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Flags:\n")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(w, "\t-%s      \t%s\n", f.Name, f.Usage)
	})
}
