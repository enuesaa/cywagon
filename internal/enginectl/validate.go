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
	e.Log.Info("******************************")
	e.Log.Info("* Sites:")
	for _, conf := range confs {
		e.Log.Info("* - %s (origin: %s)", conf.Host, conf.Origin.Url)
	}
	e.Log.Info("******************************")

	return nil
}
