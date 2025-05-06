package enginectl

import "github.com/enuesaa/cywagon/internal/libserve"

func (e *Engine) log(c *libserve.Context, format string, a ...any) {
	e.Log.Infof(c.Id+" "+format, a...)
}
