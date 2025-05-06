package service

import "github.com/enuesaa/cywagon/internal/infra"

func NewLogSrv() LogSrvInterface {
	return &LogSrv{
		Container: infra.Default,
	}
}

type LogSrvInterface interface {
	Info(text string)
	Infof(format string, a ...any)

	Debug(text string)
	Debugf(format string, a ...any)

	Err(err error)
}

type LogSrv struct {
	infra.Container
}

func (c *LogSrv) Info(text string) {
	c.Log.Info("%s", text)
}

func (c *LogSrv) Infof(format string, a ...any) {
	c.Log.Info(format, a...)
}

func (c *LogSrv) Debug(text string) {
	c.Log.Info("%s", text)
}

func (c *LogSrv) Debugf(format string, a ...any) {
	c.Log.Info(format, a...)
}

func (c *LogSrv) Err(err error) {
	c.Log.Info("%s", err.Error())
}
