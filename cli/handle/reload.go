package handle

func (h *Handler) Reload(sitename string, path string) error {
	h.Log.Info("start a new deployment..")

	if err := h.Sock.Send("deploy"); err != nil {
		return err
	}
	return nil
}
