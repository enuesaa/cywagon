package handle

import "github.com/enuesaa/cywagon/internal/enginectl"

func (h *Handler) Check(paths []string) error {
	engine := enginectl.New()

	confs, err := h.listConfs(paths)
	if err != nil {
		return err
	}
	return engine.ValidateConfs(confs)
}
