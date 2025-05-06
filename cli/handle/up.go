package handle

func (h *Handler) Up(workdir string) error {
	hclFilePaths, err := h.ConfSrv.ListHCLFiles(workdir)
	if err != nil {
		return err
	}
	config, err := h.ConfSrv.Read(hclFilePaths)
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
