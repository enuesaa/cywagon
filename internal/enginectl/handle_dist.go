package enginectl

import (
	"io/fs"
	"strings"

	"github.com/enuesaa/cywagon/internal/libserve"
)

func (e *Engine) handleDist(c *libserve.Context, dist fs.FS) *libserve.Response {
	e.logcf(c, "resolved: %s", c.Path)

	path := strings.TrimPrefix(c.Path, "/")

	f, err := dist.Open(path)
	if err != nil {
		return c.Resolve(404)
	}
	if err := c.ResBody(path, f); err != nil {
		return c.Resolve(404)
	}
	return c.Resolve(200)
}
