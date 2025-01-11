package eng

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/enuesaa/cywagon/internal/repository"
)

var ErrDownEngine = fmt.Errorf("engine down")

func Up(ctx context.Context) error {
	repos := repository.Use(ctx)

	if err := repos.Ps.CreatePidFile(); err != nil {
		return err
	}

	go Serve()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM)
	go func() {
		<-sig
		repos.Log.Info("sigterm")
		if err := Down(ctx); err != nil {
			repos.Log.Info("Error: %s", err.Error())
		}
	}()

	receiver := Receiver{}

	err := repos.Ps.ListenSocket(func(b []byte) error {
		if err := receiver.Receive(ctx, b); err != nil {
			if errors.Is(err, ErrDownEngine) {
				return err
			}
			repos.Log.Info("Error: %s", err.Error())
			return nil
		}
		return nil
	})
	return err
}
