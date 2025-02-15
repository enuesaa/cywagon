package handle

import (
	"github.com/enuesaa/cywagon/internal/enginectl"
	"github.com/enuesaa/cywagon/internal/infra"
)

func Start(confDir string) error {
	engine := enginectl.New()

	confs, err := engine.ListConfs(confDir)
	if err != nil {
		return err
	}
	for _, conf := range confs {
		if conf.Entry.Cmd != "" {
			engine.RunCmd(conf.Entry.Workdir, conf.Entry.Cmd)
		}
	}
	infra.Default.Log.Info("start serving")

	return engine.Serve(confs)
}
