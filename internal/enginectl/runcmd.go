package enginectl

import "github.com/enuesaa/cywagon/internal/infra"

type RunCmdArg struct {
	Workdir string
	Command string
}

func RunCmd(ctn infra.Container, arg RunCmdArg) {
	go runCmd(ctn, arg)
}

func runCmd(ctn infra.Container, arg RunCmdArg) {
	if err := ctn.Cmd.Start(arg.Workdir, arg.Command); err != nil {
		ctn.Log.Error(err)
	}
}
