package cli

import (
	"flag"
	"fmt"
	"io"

	"github.com/google/subcommands"
)

func FprintHelpText(w io.Writer) {
	cdr := subcommands.DefaultCommander

	fmt.Fprintf(w, "A conditionally configurable web server. Toy app\n\n")
	fmt.Fprintf(w, "Usage: \n")
	fmt.Fprintf(w, "\t%s <subcommand>\n\n", cdr.Name())

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
