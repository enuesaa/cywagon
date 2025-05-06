package enginectl

import "github.com/enuesaa/cywagon/internal/libserve"

func (e *Engine) Serve() error {
	e.Server.Port = e.config.Server.Port

	e.Log.Infof("The server started on port %d", e.Server.Port)
	for _, site := range e.config.Sites {
		e.Log.Infof("site: %s", site.Host)
	}

	e.Server.UseLogger(func(c *libserve.Context, status int, method string) {
		e.log(c, "%d %s %s", status, method, c.Path)
	})
	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		if _, ok := e.sitemap[c.Host]; !ok {
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
