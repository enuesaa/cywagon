package handle

func (h *Handler) Check(path string) error {
	config, err := h.ConfSrv.Read(path)
	if err != nil {
		return err
	}
	h.Log.Pprint(config)

	return nil
}
