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
		if _, ok := sitemap[c.Host]; !ok {
			return c.Resolve(500)
		}
		return nil
	})

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		if strings.HasSuffix(c.Path, "/") {
			c.Path = filepath.Join(c.Path, "index.html")
		}
		return nil
	})

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := sitemap[c.Host]
		for _, ifb := range site.Config.Ifs {
			if !e.matchCondStr(c.Path, ifb.Path, ifb.PathIn, ifb.PathNot, ifb.PathNotIn) {
				continue
			}
			if !e.matchCondStrMap(c.Headers, ifb.Headers, ifb.HeadersIn, ifb.HeadersNot, ifb.HeadersNotIn) {
				continue
			}
			for key, value := range ifb.Respond.Headers {
				c.ResHeader(key, value)
				return c.Resolve(500)
			}
		}
		return nil
	})
		
	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := sitemap[c.Host]
		path := strings.TrimPrefix(c.Path, "/")

		f, err := site.Dist.Open(path)
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
