package enginectl

import "github.com/enuesaa/cywagon/internal/libserve"

func (e *Engine) Serve() error {
	e.printBanner()

	e.Server.UseLogger(func(c *libserve.Context, res *libserve.Response) {
		e.log(c, "%d %s", res.GetStatus(), c.Path)
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
