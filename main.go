package main

import (
	"os"

	"github.com/enuesaa/cywagon/cli"
	"github.com/enuesaa/cywagon/internal/infra"
)

func main() {
	ctn := infra.New()

	code := cli.Run(ctn)
	os.Exit(code)
}
