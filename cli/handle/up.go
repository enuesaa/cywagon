package handle

func (h *Handler) Up(workdir string) error {
	config, err := h.ConfSrv.ReadInDir(workdir)
	if err != nil {
		return err
	}
	if err := h.Engine.Load(config, workdir); err != nil {
		return err
	}
	if err := h.Engine.Serve(); err != nil {
		return err
	}
	return nil
}
