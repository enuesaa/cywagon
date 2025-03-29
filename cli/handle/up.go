package handle

func (h *Handler) Up(paths []string) error {
	confs, err := h.ConfSrv.List(paths)
	if err != nil {
		return err
	}
	h.Engine.PrintBanner(confs)

	return h.Engine.Serve(confs)
}
