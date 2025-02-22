package enginectl

import "github.com/enuesaa/cywagon/internal/service/model"

func (e *Engine) PrintBanner(confs []model.Conf) {
	e.Log.Info("******************************")
	e.Log.Info("* Server started on localhost:3000")
	e.Log.Info("* ")
	e.Log.Info("* Sites:")
	for _, conf := range confs {
		e.Log.Info("* - %s (origin: %s)", conf.Host, conf.Origin.Url)
	}
	e.Log.Info("* ")
	e.Log.Info("******************************")
}