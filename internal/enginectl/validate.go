package enginectl

import "github.com/enuesaa/cywagon/internal/service/model"

func (e *Engine) ValidateConfs(confs []model.Conf) error {
	e.Log.Info("******************************")
	e.Log.Info("* Sites:")
	for _, conf := range confs {
		e.Log.Info("* - %s", conf.Host)
	}
	e.Log.Info("******************************")

	return nil
}
