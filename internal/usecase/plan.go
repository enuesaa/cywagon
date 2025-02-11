package usecase

import (
	"context"

	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/repository"
)

func Plan(ctx context.Context, confDir string) error {
	repos := repository.Use(ctx)

	confsrv := service.NewConfService(repos)

	files := confsrv.List(confDir)

	for _, file := range files {
		config, err := confsrv.Read(file)
		if err != nil {
			return err
		}
		repos.Log.Info("hostname: %s", config.Host)
	}
	return nil
}
