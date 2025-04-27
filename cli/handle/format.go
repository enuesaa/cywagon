package handle

func (h *Handler) Format(workdir string) error {
	hclFilePaths, err := h.ConfSrv.ListHCLFiles(workdir)
	if err != nil {
		return err
	}
	if err := h.ConfSrv.Format(hclFilePaths); err != nil {
		return err
	}
	h.Log.Info("formatted!")

	return nil
}
