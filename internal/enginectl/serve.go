package enginectl

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

type Site struct {
	Dist fs.FS
	Config model.Site
}

func (e *Engine) Serve(config model.Config) error {
	e.Server.Port = config.Server.Port

	sitemap := make(map[string]Site, 0)
	for _, site := range config.Sites {
		dist, err := e.LoadFS(site.Host, site.Dist)
		if err != nil {
			return err
		}
		sitemap[site.Host] = Site{
			Dist: dist,
			Config: site,
		}
	}

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		_, ok := sitemap[c.Host]
		if !ok {
			return c.Resolve(500)
		}
		return nil
	})

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := sitemap[c.Host]

		for _, ifblock := range site.Config.Ifs {
			if ifblock.Path != nil && *ifblock.Path == c.Path {
				for key, value := range ifblock.Respond.Headers {
					c.SetResponseHeader(key, value)
					return c.Resolve(500)
				}
			}
		}
		return nil
	})
		
	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := sitemap[c.Host]
		npath := c.Path

		if strings.HasSuffix(npath, "/") {
			npath = filepath.Join(npath, "index.html")
		}
		npath = strings.TrimPrefix(npath, "/")

		f, err := site.Dist.Open(npath)
		if err != nil {
			return c.Resolve(404)
		}
		c.SetResponseBody(npath, f)

		return c.Resolve(200)
	})

	return e.Server.Serve()
}
