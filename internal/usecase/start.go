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

	config, err := ctlconf.Read(ctx, "testdata/example.lua")
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

	return libserve.Serve(config.Entry.Host)
}
