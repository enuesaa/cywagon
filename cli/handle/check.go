package handle

func (h *Handler) Check(workdir string) error {
	hclFilePaths, err := h.ConfSrv.ListHCLFiles(workdir)
	if err != nil {
		return err
	}
	config, err := h.ConfSrv.Read(hclFilePaths)
	if err != nil {
		return err
	}
	h.Ps.Print("Configuration OK!")

	h.Ps.Print("******************************")
	h.Ps.Printf("* The server will listen on port %d", config.Server.Port)
	h.Ps.Print("* ")
	h.Ps.Print("* Sites:")
	for _, site := range config.Sites {
		h.Ps.Printf("* - %s", site.Host)
	}
	h.Ps.Print("******************************")

	return nil
}
