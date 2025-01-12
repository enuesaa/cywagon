package engctl

import (
	"context"

	"github.com/enuesaa/cywagon/internal/repository"
)

func Down(ctx context.Context) {
	repos := repository.Use(ctx)

	repos.Ps.CatchSignalStop(func() {
		repos.Log.Info("sigterm")
		if err := RemoveFiles(ctx); err != nil {
			repos.Log.Info("Error: %s", err.Error())
		}
		repos.Ps.Exit(0)
	})
}

func RemoveFiles(ctx context.Context) error {
	repos := repository.Use(ctx)
	if err := repos.Ps.DeletePidFile(); err != nil {
		return err
	}
	if err := repos.Ps.DeleteSockFile(); err != nil {
		return err
	}
	return nil
} 