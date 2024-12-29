package eng

import (
	"fmt"
	"os"
	"path/filepath"
)

func PidFile() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, "tmp/cywagon.pid"), nil
}

func CreatePidFile(pid int) error {
	path, err := PidFile()
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	content := fmt.Sprintf("%d", pid)
	if _, err := f.Write([]byte(content)); err != nil {
		return err
	}
	return nil
}

func DeletePidFile() error {
	path, err := PidFile()
	if err != nil {
		return err
	}
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}
