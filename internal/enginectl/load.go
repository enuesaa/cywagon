package enginectl

import (
	"fmt"
	"io"
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

func (e *Engine) Read(sitename string) (string, error) {
	dist, ok := e.dists[sitename]
	if !ok {
		return "", fmt.Errorf("site not found")
	}

	readme, err := dist.Open("README.md")
	if err != nil {
		return "", err
	}
	readmeBytes, err := io.ReadAll(readme)
	if err != nil {
		return "", err
	}
	return string(readmeBytes), nil
}
