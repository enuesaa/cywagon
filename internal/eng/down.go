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
	if err := repos.Ps.DeleteSockFile(); err != nil {
		return err
	}

	// TODO: subcommands.ExitStatus を返せればベスト
	os.Exit(0)

	return nil
}
