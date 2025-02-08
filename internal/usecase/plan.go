package usecase

import (
	"context"

	"github.com/enuesaa/cywagon/internal/ctlconf"
	"github.com/enuesaa/cywagon/internal/repository"
)

func Plan(ctx context.Context, confDir string) error {
	repos := repository.Use(ctx)

	files := ctlconf.List(ctx, confDir)

	for _, file := range files {
		config, err := ctlconf.Read(ctx, file)
		if err != nil {
			return err
		}
		repos.Log.Info("hostname: %s", config.Host)

		if err := config.RunHandler(ctx); err != nil {
			return err
		}
	}
	return nil
}
