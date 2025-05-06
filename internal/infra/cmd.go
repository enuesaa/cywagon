package infra

import (
	"os"
	"os/exec"
)

type CmdInterface interface {
	Start(workdir string, command string) error
}
type Cmd struct{}

func (c *Cmd) Start(workdir string, command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = workdir

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Start()
}
