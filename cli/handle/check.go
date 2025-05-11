package handle

func (h *Handler) Check(workdir string) error {
	config, err := h.ConfSrv.ReadInDir(workdir)
	if err != nil {
		return err
	}
	h.Ps.Print("Configuration OK!")

	h.Ps.Print("******************************")
	h.Ps.Print("* Sites:")
	for _, site := range config.Sites {
		h.Ps.Printf("* - %s", site.Host)
	}
	h.Ps.Print("******************************")

	return nil
}
