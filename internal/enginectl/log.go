package enginectl

import "github.com/enuesaa/cywagon/internal/libserve"

func (e *Engine) logf(format string, a ...any) {
	e.Log.Infof(format, a...)
}

func (e *Engine) logcf(c *libserve.Context, format string, a ... any) {
	e.Log.Infosf(c.Id, format, a...)
}
