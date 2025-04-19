package enginectl

import (
	"io/fs"
	"slices"

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
		site := sitemap[c.Host]

		check := func(val string, eq *string, in []string, nq *string, notin []string) bool {
			if nq != nil && *nq == val {
				return false
			}
			if len(notin) > 0 && slices.Contains(notin, val) {
				return false
			}
			if eq != nil && *eq == val {
				return true
			}
			if  len(in) > 0 && slices.Contains(in, val) {
				return true
			}
			return false
		}

		for _, ifb := range site.Config.Ifs {
			if !check(c.Path, ifb.Path, ifb.PathIn, ifb.PathNot, ifb.PathNotIn) {
				continue
			}
			for key, value := range ifb.Respond.Headers {
				c.SetResponseHeader(key, value)
				return c.Resolve(500)
			}
		}
		return nil
	})
		
	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := sitemap[c.Host]
		lpath := c.GetLookupPath()

		f, err := site.Dist.Open(lpath)
		if err != nil {
			return c.Resolve(404)
		}
		if err := c.SetResponseBody(lpath, f); err != nil {
			return c.Resolve(404)
		}
		return c.Resolve(200)
	})

	return e.Server.Serve()
}
