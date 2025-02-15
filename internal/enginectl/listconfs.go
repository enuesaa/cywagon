package enginectl

import (
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) ListConfs(confDir string) ([]model.Conf, error) {
	confsrv := service.NewConfService(e.Container)

	var confs []model.Conf

	files := confsrv.List(confDir)
	for _, file := range files {
		conf, err := confsrv.Read(file)
		if err != nil {
			return confs, err
		}
		confs = append(confs, conf)
	}
	e.Log.Info("start serving")	

	return confs, nil
}