package repository

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type PsRepositoryInterface interface {
	GetPidFilePath() (string, error)
	CreatePidFile(pid int) error
	DeletePidFile() error
	ReadPidFile() (int, error)
}

type PsRepository struct {}

func (repo *PsRepository) GetPidFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(homedir, "tmp/cywagon.pid")

	return path, nil
}


func (repo *PsRepository) CreatePidFile(pid int) error {
	path, err := repo.GetPidFilePath()
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

func (repo *PsRepository) DeletePidFile() error {
	path, err := repo.GetPidFilePath()
	if err != nil {
		return err
	}
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func (repo *PsRepository) ReadPidFile() (int, error) {
	path, err := repo.GetPidFilePath()
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
