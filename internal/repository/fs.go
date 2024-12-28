package repository

import (
	"io"
	"os"
	"path/filepath"
)

type FsRepositoryInterface interface {
	IsExist(path string) bool
	IsDir(path string) (bool, error)
	CreateDir(path string) error
	Create(path string, body []byte) error
	HomeDir() (string, error)
	WorkDir() (string, error)
	Remove(path string) error
	Read(path string) ([]byte, error)
	ListDirs(path string) ([]string, error)
	ListFiles(path string) ([]string, error)
}
type FsRepository struct{}

func (repo *FsRepository) IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (repo *FsRepository) IsDir(path string) (bool, error) {
	f, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return f.IsDir(), nil
}

func (repo *FsRepository) CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (repo *FsRepository) Create(path string, body []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(body); err != nil {
		return err
	}
	return nil
}

func (repo *FsRepository) HomeDir() (string, error) {
	return os.UserHomeDir()
}

func (repo *FsRepository) WorkDir() (string, error) {
	return os.Getwd()
}

func (repo *FsRepository) Remove(path string) error {
	return os.RemoveAll(path)
}

func (repo *FsRepository) Read(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	return io.ReadAll(f)
}

func (repo *FsRepository) ListDirs(path string) ([]string, error) {
	list := make([]string, 0)
	err := filepath.Walk(path, func(fpath string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if file.IsDir() {
			list = append(list, fpath)
		}
		return nil
	})
	return list, err
}

func (repo *FsRepository) ListFiles(path string) ([]string, error) {
	list := make([]string, 0)
	err := filepath.Walk(path, func(fpath string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if file.IsDir() {
			return nil
		}
		list = append(list, fpath)
		return nil
	})
	return list, err
}
