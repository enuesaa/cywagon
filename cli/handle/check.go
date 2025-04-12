package handle

type Config struct {
	Server Server `hcl:"server,block"`
	Sites  []Site `hcl:"site,block"`
}

type Server struct {
	Port uint `hcl:"port"`
}

type Site struct {
	Name string `hcl:"name,label"`
	Host string `hcl:"host"`
	Dist string `hcl:"dist"`
}

func (h *Handler) Check(path string) error {
	config, err := h.ConfSrv.Read(path)
	if err != nil {
		return err
	}
	h.Log.Info("conf %+v", config)

	return nil
}
