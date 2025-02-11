package infra

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
}

type LogInterface interface {
	Error(err error)
	Info(format string, a ...any)
}

type Log struct{}

func (i *Log) print(text string) {
	log.Printf("%s\n", text)
}

func (i *Log) Info(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	i.print(text)
}

func (i *Log) Error(err error) {
	text := fmt.Sprintf("Error: %s", err.Error())
	fmt.Fprintf(os.Stderr, "%s\n", text)
}
