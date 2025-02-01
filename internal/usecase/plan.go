package usecase

import (
	"context"

	"github.com/enuesaa/cywagon/internal/ctlconf"
	"github.com/enuesaa/cywagon/internal/repository"
)

func Plan(ctx context.Context, confDir string) error {
	repos := repository.Use(ctx)

	config, err := ctlconf.Parse(ctx, "testdata/sites-enabled/example.lua")
	if err != nil {
		return err
	}
	repos.Log.Print("hostname: %s\n", config.Hostname)
	repos.Log.Print("port: %d\n", config.Port)

	if err := config.RunHandler(); err != nil {
		return err
	}

	return nil
}
