package handle

import "github.com/enuesaa/cywagon/internal/enginectl"

func (h *Handler) Up(confDir string) error {
	engine := enginectl.New()

	confs, err := engine.ListConfs(confDir)
	if err != nil {
		return err
	}
	if err := engine.StartUp(confs); err != nil {
		return err
	}
	if err := engine.StartHealthCheck(confs); err != nil {
		return err
	}
	h.Log.Info("start serving")

	return engine.Serve(confs)
}
