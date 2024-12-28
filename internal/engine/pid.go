package engine

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
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

func ReadPidFile() (int, error) {
	path, err := PidFile()
	if err != nil {
		return -1, err
	}

	f, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return -1, err
	}
	content := string(bytes)

	pid, err := strconv.Atoi(content)
	if err != nil {
		return -1, err
	}
	return pid, nil
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
