package handle

import (
	"github.com/enuesaa/cywagon/internal/enginectl"
	"github.com/enuesaa/cywagon/internal/infra"
)

func Up(confDir string) error {
	engine := enginectl.New()

	confs, err := engine.ListConfs(confDir)
	if err != nil {
		return err
	}
	if err := engine.StartUp(confs); err != nil {
		return err
	}
	infra.Default.Log.Info("start serving")

	if err := engine.CheckHealth(); err != nil {
		return err
	}

	return engine.Serve(confs)
}
