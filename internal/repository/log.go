package repository

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(0)
}

type LogRepositoryInterface interface {
	Error(err error)
	Info(format string, a ...any)
}

type LogRepository struct{}

func (repo *LogRepository) print(text string) {
	log.Printf("%s\n", text)
}

func (repo *LogRepository) Info(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	repo.print(text)
}

func (repo *LogRepository) Error(err error) {
	text := fmt.Sprintf("Error: %s", err.Error())
	repo.print(text)
}
