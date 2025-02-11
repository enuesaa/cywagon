package main

import (
	"os"

	"github.com/enuesaa/cywagon/internal/cli"
	"github.com/enuesaa/cywagon/internal/repository"
)

func main() {
	repos := repository.New()

	code := cli.Run(repos)
	os.Exit(code)
}
