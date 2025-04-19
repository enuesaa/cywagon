package enginectl

import (
	"io/fs"
	"os"
)

func (e *Engine) LoadFS(sitename string, path string) (fs.FS, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	return os.DirFS(path), nil
}
