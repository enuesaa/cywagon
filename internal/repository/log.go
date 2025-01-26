package repository

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func init() {
	log.SetFlags(0)
}

type LogRepositoryInterface interface {
	PrintErr(err error)
	Print(format string, a ...any)
	Info(format string, a ...any) error
}

type LogRepository struct{}

func (repo *LogRepository) PrintErr(err error) {
	log.Printf("Error: %s\n", err.Error())
}

func (repo *LogRepository) Print(format string, a ...any) {
	log.Printf(format, a...)
}

func (repo *LogRepository) Info(format string, a ...any) error {
	text := fmt.Sprintf(format, a...)
	text = fmt.Sprintf("%s: %s\n", time.Now().Local().Format(time.RFC3339), text)

	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(homedir, "tmp/cywagon.log")

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(text); err != nil {
		return err
	}
	return nil
}
