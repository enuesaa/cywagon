package infra

import (
	"io"
	"os"
	"path/filepath"
)

// To generate mock file, run following command:
//   mockgen -source=fs.go -destination=fs_mock.go -package=infra

type FsInterface interface {
	IsExist(path string) bool
	IsFile(path string) bool
	CreateDir(path string) error
	Create(path string, body []byte) error
	HomeDir() (string, error)
	WorkDir() (string, error)
	Remove(path string) error
	Read(path string) ([]byte, error)
	ListDirs(path string) ([]string, error)
	ListFiles(path string) ([]string, error)
}
type Fs struct{}

func (i *Fs) IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (i *Fs) IsFile(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

func (i *Fs) CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (i *Fs) Create(path string, body []byte) error {
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

func (i *Fs) HomeDir() (string, error) {
	return os.UserHomeDir()
}

func (i *Fs) WorkDir() (string, error) {
	return os.Getwd()
}

func (i *Fs) Remove(path string) error {
	return os.RemoveAll(path)
}

func (i *Fs) Read(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	return io.ReadAll(f)
}

func (i *Fs) ListDirs(path string) ([]string, error) {
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

func (i *Fs) ListFiles(path string) ([]string, error) {
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
