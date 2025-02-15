package enginectl

import (
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) ListConfs(dir string) ([]model.Conf, error) {
	var confs []model.Conf

	confsrv := service.NewConfService()
	files := confsrv.List(dir)

	for _, file := range files {
		conf, err := confsrv.Read(file)
		if err != nil {
			return nil, err
		}
		confs = append(confs, conf)
	}

	return confs, nil
}
