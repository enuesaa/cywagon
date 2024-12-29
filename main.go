package main

import (
	"os"

	"github.com/enuesaa/cywagon/internal/cli"
	"github.com/enuesaa/cywagon/internal/repository"
)

func main() {
	ctx := repository.NewContext()

	code := cli.Run(ctx)
	os.Exit(code)
}
