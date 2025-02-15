package main

import (
	"github.com/enuesaa/cywagon/cli"
	"github.com/enuesaa/cywagon/internal/infra"
)

func main() {
	code := cli.Run()
	infra.I.Ps.Exit(code)
}
