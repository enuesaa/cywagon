package enginectl

func (e *Engine) RunCmd(workdir string, command string) {
	go e.runCmd(workdir, command)
}

func (e *Engine) runCmd(workdir string, command string) {
	if err := e.Cmd.Start(workdir, command); err != nil {
		e.Log.Error(err)
	}
}
