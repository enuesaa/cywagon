package enginectl

func (e *Engine) RunCmd(workdir string, command string) {
	fn := func() {
		if err := e.Cmd.Start(workdir, command); err != nil {
			e.Log.Error(err)
		}
	}
	go fn()
}
