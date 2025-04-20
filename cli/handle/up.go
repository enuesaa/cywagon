package handle

func (h *Handler) Up(workdir string) error {
	config, err := h.ConfSrv.ReadInWorkdir(workdir)
	if err != nil {
		return err
	}
	if err := h.Engine.Serve(config); err != nil {
		return err
	}
	if err := h.Engine.Close(); err != nil {
		return err
	}
	return nil
}
