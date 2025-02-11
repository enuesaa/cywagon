package handle

import (
	"github.com/enuesaa/cywagon/internal/enginectl"
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func Start(ctn infra.Container, confDir string) error {
	confsrv := service.NewConfService(ctn)

	var confs []model.Conf

	files := confsrv.List(confDir)
	for _, file := range files {
		conf, err := confsrv.Read(file)
		if err != nil {
			return err
		}
		if conf.Entry.Cmd != "" {
			enginectl.RunCmd(ctn, enginectl.RunCmdArg{
				Workdir: conf.Entry.Workdir,
				Command: conf.Entry.Cmd,
			})
		}
		confs = append(confs, conf)
	}
	ctn.Log.Info("start serving")

	sites := make([]libserve.ServeOptsSite, 0)
	for _, conf := range confs {
		sites = append(sites, libserve.ServeOptsSite{
			Host:      conf.Host,
			OriginUrl: conf.Entry.Host,
			Handler:   conf.RunHandler,
		})
	}
	serveOpts := libserve.ServeOpts{
		Port:  3000,
		Sites: sites,
	}

	return libserve.Serve(serveOpts)
}
