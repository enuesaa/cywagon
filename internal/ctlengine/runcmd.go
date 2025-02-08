package ctlengine

import (
	"context"

	"github.com/enuesaa/cywagon/internal/repository"
)

type RunCmdArg struct {
	Workdir string
	Command string
}

func RunCmd(ctx context.Context, arg RunCmdArg) {
	go runCmd(ctx, arg)
}

func runCmd(ctx context.Context, arg RunCmdArg) {
	repos := repository.Use(ctx)

	if err := repos.Cmd.Start(arg.Workdir, arg.Command); err != nil {
		repos.Log.Error(err)
	}
}
