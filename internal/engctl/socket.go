package engctl

import (
	"os"
	"path/filepath"
)

func Socket() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, "tmp/cywagon.sock"), nil
}
