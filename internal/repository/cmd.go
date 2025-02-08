package repository

import (
	"os"
	"os/exec"
)

type CmdRepositoryInterface interface {
	Start(workdir string, command string) error
}
type CmdRepository struct{}

func (c *CmdRepository) Start(workdir string, command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = workdir

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Start()
}
