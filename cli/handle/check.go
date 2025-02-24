package handle

func (h *Handler) Check(paths []string) error {
	confs, err := h.ConfSrv.List(paths)
	if err != nil {
		return err
	}
	return h.Engine.ValidateConfs(confs)
}
