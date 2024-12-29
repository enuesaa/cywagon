package engctl

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func PidFile() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, "tmp/cywagon.pid"), nil
}

func ReadPidFile() (int, error) {
	path, err := PidFile()
	if err != nil {
		return -1, err
	}

	f, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return -1, err
	}
	content := string(bytes)

	pid, err := strconv.Atoi(content)
	if err != nil {
		return -1, err
	}
	return pid, nil
}
