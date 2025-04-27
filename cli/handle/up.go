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
	if err := h.Engine.Serve(config, workdir); err != nil {
		return err
	}
	if err := h.Engine.Close(); err != nil {
		return err
	}
	return nil
}
