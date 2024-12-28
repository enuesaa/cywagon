package internal

import (
	"fmt"
	"os/exec"
	"syscall"
)

func Up() error {
	cmd := exec.Command("cywagon", "engine-start")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	fmt.Printf("pid: %d\n", cmd.Process.Pid)

	return nil
}