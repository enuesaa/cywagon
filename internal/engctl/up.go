package engctl

import (
	"context"
	"fmt"
	"os/exec"
	"syscall"
)

func Up(ctx context.Context) error {
	cmd := exec.Command("cywagon", "up", "--foreground")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	fmt.Printf("pid: %d\n", cmd.Process.Pid)

	return nil
}
