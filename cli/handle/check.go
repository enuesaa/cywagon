package handle

func (h *Handler) Check(workdir string) error {
	config, err := h.ConfSrv.Read(workdir)
	if err != nil {
		return err
	}
	h.Log.Info("Configuration OK!")

	h.Log.Info("******************************")
	h.Log.Info("* The server will listen on port %d", config.Server.Port)
	h.Log.Info("* ")
	h.Log.Info("* Sites:")
	for _, site := range config.Sites {
		h.Log.Info("* - %s", site.Host)
	}
	h.Log.Info("******************************")

	return nil
}
