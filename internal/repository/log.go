package repository

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(0)
}

type LogRepositoryInterface interface {
	Printf(format string, a ...any)
	Fatalf(format string, a ...any)
	GetAll() string
}

type LogRepository struct {
	out string
}

func (repo *LogRepository) Printf(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	repo.out += text
	fmt.Print(text)
}

func (repo *LogRepository) Fatalf(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	repo.out += text
	// todo print out all logs here.
	log.Fatal(text)
}

func (repo *LogRepository) GetAll() string {
	return repo.out
}
