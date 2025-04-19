package enginectl

import (
	"io/fs"
	"os"
)

type Dists map[string]fs.FS

func (e *Engine) LoadFS(sitename string, path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	e.dists[sitename] = os.DirFS(path)

	return nil
}
