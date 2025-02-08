package usecase

import (
	"context"

	"github.com/enuesaa/cywagon/internal/ctlconf"
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
		repos.Log.Info("%+v", conf)

		if conf.Entry.Cmd != "" {
			go func() {
				if err := repos.Cmd.Start(conf.Entry.Workdir, conf.Entry.Cmd); err != nil {
					repos.Log.Error(err)
				}
			}()
		}
		confs = append(confs, conf)
	}
	repos.Log.Info("start serving")

	return libserve.Serve(confs)
}
