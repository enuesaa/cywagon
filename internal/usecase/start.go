package usecase

import (
	"context"
	"net/http"

	"github.com/enuesaa/cywagon/internal/ctlconf"
	"github.com/enuesaa/cywagon/internal/ctlengine"
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/repository"
)

func Start(ctx context.Context, confDir string) error {
	repos := repository.Use(ctx)

	var confs []ctlconf.Conf

	files := ctlconf.List(ctx, confDir)
	for _, file := range files {
		conf, err := ctlconf.Read(ctx, file)
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
			Handler: func(req *http.Request, fn libserve.FnNext, res *http.Response) error {
				return conf.RunHandler(req, fn, res)
			},
		})
	}
	serveOpts := libserve.ServeOpts{
		Port:  3000,
		Sites: sites,
	}

	return libserve.Serve(serveOpts)
}
