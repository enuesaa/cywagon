package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func LogFile() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, "tmp/cywagon.log"), nil
}

func Log(message string) error {
	message = fmt.Sprintf("%s: %s\n", time.Now().Local().Format(time.RFC3339), message)

	path, err := LogFile()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(message); err != nil {
		return err
	}
	return nil
}
