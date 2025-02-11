package usecase

import (
	"context"

	"github.com/enuesaa/cywagon/internal/ctlengine"
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func Start(ctx context.Context, confDir string) error {
	repos := repository.Use(ctx)

	confsrv := service.NewConfService(repos)

	var confs []model.Conf

	files := confsrv.List(confDir)
	for _, file := range files {
		conf, err := confsrv.Read(file)
		if err != nil {
			return err
		}
		if conf.Entry.Cmd != "" {
			ctlengine.RunCmd(ctx, ctlengine.RunCmdArg{
				Workdir: conf.Entry.Workdir,
				Command: conf.Entry.Cmd,
			})
		}
		confs = append(confs, conf)
	}
	repos.Log.Info("start serving")

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
