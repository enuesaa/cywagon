package enginectl

func (e *Engine) printBanner() {
	e.Log.Infof("Load completed. The server will start on port %d", e.Server.Port)
}
