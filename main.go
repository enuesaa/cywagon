package main

import (
	"os"

	"github.com/enuesaa/cywagon/cli"
	"github.com/enuesaa/cywagon/internal/infra"
)

func main() {
	container := infra.New()

	code := cli.Run(container)
	os.Exit(code)
}
