package handle

import (
	"github.com/enuesaa/cywagon/internal/enginectl"
	"github.com/enuesaa/cywagon/internal/infra"
)

func Start(ctn infra.Container, confDir string) error {
	confs, err := enginectl.ListConfs(ctn, confDir)
	if err != nil {
		return err
	}
	for _, conf := range confs {
		if conf.Entry.Cmd != "" {
			enginectl.RunCmd(ctn, conf.Entry.Workdir, conf.Entry.Cmd)
		}
	}

	return enginectl.Serve(ctn, confs)
}
