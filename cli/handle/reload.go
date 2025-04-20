package handle

func (h *Handler) Reload(workdir string) error {
	if err := h.Sock.Send("reload"); err != nil {
		return err
	}
	return nil
}
