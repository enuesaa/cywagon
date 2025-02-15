package enginectl

import (
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func ListConfs(ctn infra.Container, confDir string) ([]model.Conf, error) {
	confsrv := service.NewConfService(ctn)

	var confs []model.Conf

	files := confsrv.List(confDir)
	for _, file := range files {
		conf, err := confsrv.Read(file)
		if err != nil {
			return confs, err
		}
		confs = append(confs, conf)
	}
	ctn.Log.Info("start serving")	

	return confs, nil
}