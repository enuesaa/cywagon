package handle

func (h *Handler) Publish(sitename string, deploymentId string) error {
	h.Log.Info("start publishing..")

	if err := h.Sock.Send("publish"); err != nil {
		return err
	}
	return nil
}
