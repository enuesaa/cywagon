package handle

import "github.com/enuesaa/cywagon/internal/enginectl"

func (h *Handler) Check(confDir string) error {
	engine := enginectl.New()

	confs, err := engine.ListConfs(confDir)
	if err != nil {
		return err
	}

	return engine.ValidateConfs(confs)
}
