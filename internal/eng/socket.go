package eng

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

func DeleteSockFile() error {
	path, err := Socket()
	if err != nil {
		return err
	}
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}
