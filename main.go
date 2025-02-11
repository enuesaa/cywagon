package main

import (
	"github.com/enuesaa/cywagon/cli"
	"github.com/enuesaa/cywagon/internal/infra"
)

func main() {
	ctn := infra.New()

	code := cli.Run(ctn)
	ctn.Ps.Exit(code)
}
