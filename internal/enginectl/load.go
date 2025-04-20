package enginectl

import (
	"io/fs"
	"os"
	"path/filepath"
)

func (e *Engine) LoadFS(workdir string, distpath string) (fs.FS, error) {
	path := filepath.Join(workdir, distpath)
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	return os.DirFS(path), nil
}
