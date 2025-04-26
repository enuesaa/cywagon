package enginectl

import (
	"io/fs"
	"strings"

	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

type Site struct {
	Dist fs.FS
	Config model.Site
}

func (e *Engine) loadSites(config model.Config) map[string]model.Site {
	sitemap := make(map[string]model.Site, 0)
	for _, site := range config.Sites {
		sitemap[site.Host] = site
	}
	return sitemap
}

func (e *Engine) loadDists(config model.Config, workdir string) (map[string]fs.FS, error) {
	distmap := make(map[string]fs.FS)
	for _, site := range config.Sites {
		dist, err := e.LoadFS(workdir, site.Dist)
		if err != nil {
			return nil, err
		}
		distmap[site.Dist] = dist

		for _, ifb := range site.Ifs {
			if ifb.Respond != nil && ifb.Respond.Dist != nil {
				distpath := *ifb.Respond.Dist
				dist, err := e.LoadFS(workdir, distpath)
				if err != nil {
					return nil, err
				}
				distmap[distpath] = dist
			}
		}
	}
	return distmap, nil
}

func (e *Engine) Serve(config model.Config, workdir string) error {
	e.Server.Port = config.Server.Port

	sitemap := e.loadSites(config)
	distmap, err := e.loadDists(config, workdir)
	if err != nil {
		return err
	}

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		if _, ok := sitemap[c.Host]; !ok {
			return c.Resolve(500)
		}
		return nil
	})

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := sitemap[c.Host]
		for name, value := range site.Headers {
			c.ResHeader(name, value)
		}
		return nil
	})

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := sitemap[c.Host]
		for _, ifb := range site.Ifs {
			if e.shouldCheckCondStr(ifb.Path, ifb.PathIn, ifb.PathNot, ifb.PathNotIn) {
				if !e.matchCondPath(c.Path, ifb.Path, ifb.PathIn, ifb.PathNot, ifb.PathNotIn) {
					continue
				}
			}
			if e.shouldCheckCondStrMap(ifb.Headers, ifb.HeadersIn, ifb.HeadersNot, ifb.HeadersNotIn) {
				if !e.matchCondStrMap(c.Headers, ifb.Headers, ifb.HeadersIn, ifb.HeadersNot, ifb.HeadersNotIn) {
					continue
				}
			}
			if ifb.Rewrite != nil {
				if ifb.Rewrite.Path != nil {
					c.Path = *ifb.Rewrite.Path
				}
			}
			if ifb.Respond != nil {
				for key, value := range ifb.Respond.Headers {
					c.ResHeader(key, value)
				}
				if ifb.Respond.Body != nil {
					c.ResBody(c.Path, strings.NewReader(*ifb.Respond.Body))
				}
				if ifb.Respond.Status != nil {
					c.ResStatusPrefer(*ifb.Respond.Status)
				}
				if ifb.Respond.Dist != nil {
					distpath := *ifb.Respond.Dist
					dist := distmap[distpath]
					path := strings.TrimPrefix(c.Path, "/")
		
					f, err := dist.Open(path)
					if err != nil {
						return c.Resolve(404)
					}
					if err := c.ResBody(path, f); err != nil {
						return c.Resolve(404)
					}
				}
				return c.Resolve(200)
			}
		}
		return nil
	})

	e.Server.Use(func(c *libserve.Context) *libserve.Response {
		site := sitemap[c.Host]
		dist := distmap[site.Dist]
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
