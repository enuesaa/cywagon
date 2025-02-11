package enginectl

import (
	"context"

	"github.com/enuesaa/cywagon/internal/infra"
)

type RunCmdArg struct {
	Workdir string
	Command string
}

func RunCmd(ctx context.Context, arg RunCmdArg) {
	go runCmd(ctx, arg)
}

func runCmd(ctx context.Context, arg RunCmdArg) {
	repos := infra.Use(ctx)

	if err := repos.Cmd.Start(arg.Workdir, arg.Command); err != nil {
		repos.Log.Error(err)
	}
}
