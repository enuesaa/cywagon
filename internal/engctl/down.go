package engctl

import (
	"context"

	"github.com/enuesaa/cywagon/internal/repository"
)

func Down(ctx context.Context) error {
	repos := repository.Use(ctx)

	pid, err := repos.Ps.ReadPidFile()
	if err != nil {
		return err
	}
	return repos.Ps.SendSigTerm(pid)
}
