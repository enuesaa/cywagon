package enginectl

import (
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) handleHeader(c *libserve.Context, site model.Site) *libserve.Response {
	for name, value := range site.Headers {
		c.ResHeader(name, value)
	}
	return nil
}
