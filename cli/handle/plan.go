package handle

import (
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/infra"
)

func Plan(ctn infra.Container, confDir string) error {
	confsrv := service.NewConfService(ctn)

	files := confsrv.List(confDir)

	for _, file := range files {
		config, err := confsrv.Read(file)
		if err != nil {
			return err
		}
		ctn.Log.Info("hostname: %s", config.Host)
	}
	return nil
}
