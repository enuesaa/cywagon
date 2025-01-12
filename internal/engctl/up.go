package engctl

import (
	"context"
	"fmt"

	"github.com/enuesaa/cywagon/internal/repository"
)

var ErrDownEngine = fmt.Errorf("engine down")

func Up(ctx context.Context) error {
	repos := repository.Use(ctx)

	if err := repos.Ps.CreatePidFile(); err != nil {
		return err
	}

	go Serve(ctx)
	go Down(ctx)

	err := repos.Ps.ListenSocket(func(b []byte) error {
		return nil
	})
	return err
}
