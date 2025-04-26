package enginectl

import (
	"io/fs"
	"strings"

	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

type Site struct {
	Dist   fs.FS
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

func (e *Engine) loadLogics(config model.Config) (map[string]model.Logic) {
	logicmap := make(map[string]model.Logic)
	for _, logic := range config.Logics {
		logicmap[logic.Name] = logic
	}
	return logicmap
}

func (e *Engine) Serve(config model.Config, workdir string) error {
	e.Server.Port = config.Server.Port

	sitemap := e.loadSites(config)
	distmap, err := e.loadDists(config, workdir)
	if err != nil {
		return err
	}
	logicmap := e.loadLogics(config)

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
		return e.handleIfBlocks(c, site.Ifs, distmap, logicmap)
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
