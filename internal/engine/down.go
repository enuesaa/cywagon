package engine

import (
	"fmt"
	"os"
	"syscall"
)

func Down() error {
	pid, err := ReadPidFile()
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

	if err := DeletePidFile(); err != nil {
		return err
	}
	return nil
}
