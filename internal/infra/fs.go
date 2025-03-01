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
	Read(path string) ([]byte, error)
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

func (i *Fs) Read(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	return io.ReadAll(f)
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
