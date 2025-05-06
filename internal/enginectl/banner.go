package enginectl

func (e *Engine) printBanner() {
	e.Log.Infof("The server started on port %d", e.Server.Port)
}
