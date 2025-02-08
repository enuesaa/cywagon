package repository

import "log"

func init() {
	log.SetFlags(0)
}

type LogRepositoryInterface interface {
	Error(err error)
	Info(format string, a ...any) error
}

type LogRepository struct{}

func (repo *LogRepository) Error(err error) {
	log.Printf("Error: %s\n", err.Error())
}

func (repo *LogRepository) Info(format string, a ...any) error {
	log.Printf(format, a...)
	return nil
}
