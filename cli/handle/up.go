package handle

import "github.com/enuesaa/cywagon/internal/enginectl"

func (h *Handler) Up(paths []string) error {
	engine := enginectl.New(h.Container)

	confs, err := h.listConfs(paths)
	if err != nil {
		return err
	}

	h.Log.Info("Start up sites..")
	if err := engine.StartUp(confs); err != nil {
		return err
	}

	h.Log.Info("Start health check..")
	if err := engine.StartHealthCheck(confs); err != nil {
		return err
	}
	engine.PrintBanner(confs)

	return engine.Serve(confs)
}
