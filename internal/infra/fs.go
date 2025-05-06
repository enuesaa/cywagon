package infra

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type FsInterface interface {
	IsExist(path string) bool
	IsFile(path string) bool
	Read(path string) ([]byte, error)
	Create(path string, body []byte) error
	ListFiles(path string) ([]string, error)
	DirFS(path string) (fs.FS, error)
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

func (i *Fs) Read(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	return io.ReadAll(f)
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

func (i *Fs) DirFS(path string) (fs.FS, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	return os.DirFS(path), nil
}
