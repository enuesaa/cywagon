package eng

import (
	"context"
	"os"

	"github.com/enuesaa/cywagon/internal/repository"
)

func Up(ctx context.Context) error {
	repos := repository.Use(ctx)

	pid := os.Getegid()
	if err := repos.Ps.CreatePidFile(pid); err != nil {
		return err
	}

	err := repos.Ps.ListenSocket(func(b []byte) error {
		// TODO: handle err carefully.
		if err := Receive(ctx, b); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		repos.Log.Info("Error: %s", err.Error())
		return err
	}
	return nil
}
