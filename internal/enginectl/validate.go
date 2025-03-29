package enginectl

import "github.com/enuesaa/cywagon/internal/service/model"

func (e *Engine) ValidateConfs(confs []model.Conf) error {
	for _, conf := range confs {
		if err := e.ConfSrv.Validate(conf); err != nil {
			return err
		}
	}
	e.Log.Info("******************************")
	e.Log.Info("* Sites:")
	for _, conf := range confs {
		e.Log.Info("* - %s", conf.Host)
	}
	e.Log.Info("******************************")

	return nil
}
