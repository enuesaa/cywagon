package enginectl

import (
	"strings"

	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) Serve(config model.Config, workdir string) error {
	e.Server.Port = config.Server.Port

	e.loadSites(config)
	e.loadLogics(config)
	if err := e.loadDists(config, workdir); err != nil {
		return err
	}

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		if _, ok := e.sitemap[c.Host]; !ok {
			return c.Resolve(500)
		}
		return nil
	})

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := e.sitemap[c.Host]
		for name, value := range site.Headers {
			c.ResHeader(name, value)
		}
		return nil
	})

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := e.sitemap[c.Host]
		return e.handleIfBlocks(c, site.Ifs)
	})

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := e.sitemap[c.Host]
		dist := e.distmap[site.Dist]
		path := strings.TrimPrefix(c.Path, "/")

		f, err := dist.Open(path)
		if err != nil {
			return c.Resolve(404)
		}
		if err := c.ResBody(path, f); err != nil {
			return c.Resolve(404)
		}
		return c.Resolve(200)
	})

	return e.Server.Serve()
}
