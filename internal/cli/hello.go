package cli

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type helloCmd struct {}

func (c *helloCmd) Name() string {
	return "hello"
}

func (c *helloCmd) Synopsis() string {
	return "hello"
}

func (c *helloCmd) Usage() string {
	return `hello:
	Print hello.
  `
}

func (c *helloCmd) SetFlags(f *flag.FlagSet) {}

func (c *helloCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return subcommands.ExitSuccess
}
