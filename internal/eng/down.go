package eng

import (
	"context"
	"os"

	"github.com/enuesaa/cywagon/internal/repository"
)

func Down(ctx context.Context) error {
	repos := repository.Use(ctx)

	if err := repos.Ps.DeletePidFile(); err != nil {
		return err
	}
	if err := DeleteSockFile(); err != nil {
		return err
	}

	os.Exit(0)

	return nil
}
