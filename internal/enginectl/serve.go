package enginectl

import "github.com/enuesaa/cywagon/internal/libserve"

func (e *Engine) Serve() error {
	e.Server.OnResponse = func(c *libserve.Context, status int, method string) {
		e.logcf(c, "%d %s %s (as %s)", status, method, c.GetOriginalPath(), c.Path)
	}
	e.Server.OnError = func(c *libserve.Context, err error) {
		e.logcf(c, "err: %s", err.Error())
	}

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site, ok := e.sitemap[c.Host]
		if !ok {
			return c.Resolve(500)
		}
		if site.Port != c.Port {
			return c.Resolve(500)
		}
		return nil
	})
	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := e.sitemap[c.Host]
		return e.handleHeader(c, site)
	})
	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := e.sitemap[c.Host]
		return e.handleIfBlocks(c, site.Ifs)
	})
	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := e.sitemap[c.Host]
		dist := e.distmap[site.Dist]
		return e.handleDist(c, dist)
	})

	return e.Server.Serve()
}
