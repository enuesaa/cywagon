package repository

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

type PsRepositoryInterface interface {
	GetPidFilePath() (string, error)
	CreatePidFile() error
	DeletePidFile() error
	ReadPidFile() (int, error)
	GetSockPath() (string, error)
	DeleteSockFile() error
	SendThroughSocket(data []byte) error
	ListenSocket(callback func(b []byte) error) error
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


func (repo *PsRepository) CreatePidFile() error {
	path, err := repo.GetPidFilePath()
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	pid := os.Getpid()
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

func (repo *PsRepository) GetSockPath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(homedir, "tmp/cywagon.sock")

	return path, nil
}

func (repo *PsRepository) DeleteSockFile() error {
	path, err := repo.GetSockPath()
	if err != nil {
		return err
	}
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func (repo *PsRepository) SendThroughSocket(data []byte) error {
	socket, err := repo.GetSockPath()
	if err != nil {
		return err
	}
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// Listen and wait until err fallback
func (repo *PsRepository) ListenSocket(callback func(b []byte) error) error {
	socket, err := repo.GetSockPath()
	if err != nil {
		return err
	}
	listener, err := net.Listen("unix", socket)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		defer conn.Close()

		bytes, err := io.ReadAll(conn)
		if err != nil {
			return err
		}
		if err := callback(bytes); err != nil {
			return err
		}
	}
}
