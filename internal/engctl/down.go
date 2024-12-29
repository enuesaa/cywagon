package engctl

import (
	"context"
	"fmt"
	"os"
	"syscall"

	"github.com/enuesaa/cywagon/internal/repository"
)

func Down(ctx context.Context) error {
	repos := repository.Use(ctx)

	pid, err := repos.Ps.ReadPidFile()
	if err != nil {
		return err
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	if err := process.Signal(syscall.SIGTERM); err != nil {
		return err
	}
	fmt.Println("send sigterm")
	return nil
}
