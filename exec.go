package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func execCmd() error {
	cmd := exec.Command("sleep", "10")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	fmt.Printf("pid: %d\n", cmd.Process.Pid)

	return fmt.Errorf("end")
}
