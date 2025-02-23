package enginectl

import (
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) ValidateConfs(confs []model.Conf) error {
	confsrv := service.NewConfService()

	for _, conf := range confs {
		if err := confsrv.Validate(conf); err != nil {
			return err
		}
	}
	return nil
}
