package handle

func (h *Handler) Format(workdir string) error {
	if err := h.ConfSrv.Format(workdir); err != nil {
		return err
	}
	h.Log.Info("formatted!")

	return nil
}
