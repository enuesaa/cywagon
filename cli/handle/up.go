package handle

func (h *Handler) Up(path string) error {
	config, err := h.ConfSrv.Read(path)
	if err != nil {
		return err
	}
	h.Log.Info("conf %+v", config)

	if err := h.Engine.Serve(config); err != nil {
		return err
	}
	if err := h.Engine.Close(); err != nil {
		return err
	}
	return nil
}
