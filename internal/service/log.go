package service

import (
	"fmt"
	"io"
	"os"

	"github.com/enuesaa/cywagon/internal/infra"
)

func NewLogSrv() LogSrvInterface {
	return &LogSrv{
		Container: infra.Default,
		writer: os.Stdout,
	}
}

type LogSrvInterface interface {
	SetLogFile(path string) error

	Info(text string)
	Infof(format string, a ...any)

	Debug(text string)
	Debugf(format string, a ...any)

	Err(err error)
}

type LogSrv struct {
	infra.Container

	writer io.Writer
}

func (c *LogSrv) SetLogFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	c.writer = f

	return nil
}

func (c *LogSrv) Info(text string) {
	fmt.Fprintf(c.writer, "[info] %s\n", text)
}

func (c *LogSrv) Infof(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	c.Info(text)
}

func (c *LogSrv) Debug(text string) {
	fmt.Fprintf(c.writer, "[debug] %s\n", text)
}

func (c *LogSrv) Debugf(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	c.Debug(text)
}

func (c *LogSrv) Err(err error) {
	fmt.Fprintf(c.writer, "[err] %s\n", err.Error())
}
