package main

import (
	"github.com/enuesaa/cywagon/cli"
	"github.com/enuesaa/cywagon/internal/infra"
)

func main() {
	container := infra.New()

	code := cli.Run(container)
	container.Ps.Exit(code)
}
