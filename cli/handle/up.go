package handle

func (h *Handler) Up(paths []string) error {
	confs, err := h.ConfSrv.List(paths)
	if err != nil {
		return err
	}

	h.Log.Info("Start up sites..")
	if err := h.Engine.StartUp(confs); err != nil {
		return err
	}

	h.Log.Info("Start health check..")
	if err := h.Engine.StartHealthCheck(confs); err != nil {
		return err
	}
	h.Engine.PrintBanner(confs)

	return h.Engine.Serve(confs)
}
