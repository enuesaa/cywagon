package usecase

import (
	"context"
	"time"

	"github.com/enuesaa/cywagon/internal/ctlconf"
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/repository"
)

func Start(ctx context.Context, confDir string) error {
	repos := repository.Use(ctx)

	files := ctlconf.List(ctx, confDir)

	for _, file := range files {
		config, err := ctlconf.Read(ctx, file)
		if err != nil {
			return err
		}
		repos.Log.Info("%+v", config)

		go func() {
			if err := repos.Cmd.Start(config.Entry.Workdir, config.Entry.Cmd); err != nil {
				repos.Log.Error(err)
			}
		}()
		time.Sleep(time.Duration(config.Entry.WaitForHealthy) * time.Second)
		repos.Log.Info("start serving")

		if err := libserve.Serve(config.Entry.Host); err != nil {
			return err
		}
	}
	return nil
}
