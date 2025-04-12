package enginectl

import "github.com/enuesaa/cywagon/internal/service/model"

func (e *Engine) PrintBanner(confs []model.Config) {
	e.Log.Info("******************************")
	e.Log.Info("* Server started on localhost:3000")
	e.Log.Info("* ")
	e.Log.Info("* Sites:")
	// for _, conf := range confs {
	// 	e.Log.Info("* - %s", conf.Host)
	// }
	e.Log.Info("******************************")
}
