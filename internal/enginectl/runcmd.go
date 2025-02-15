package enginectl

import "github.com/enuesaa/cywagon/internal/infra"

func RunCmd(ctn infra.Container, workdir string, command string) {
	go runCmd(ctn, workdir, command)
}

func runCmd(ctn infra.Container, workdir string, command string) {
	if err := ctn.Cmd.Start(workdir, command); err != nil {
		ctn.Log.Error(err)
	}
}
