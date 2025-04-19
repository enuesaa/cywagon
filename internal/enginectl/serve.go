package enginectl

import (
	"io/fs"
	"net/http"
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

	e.Server.Use(func(res *libserve.Response, req *http.Request) {
		_, ok := sitemap[req.Host]
		if !ok {
			res.Write(500, "", nil)
			return
		}
	})

	e.Server.Use(func(res *libserve.Response, req *http.Request) {
		site := sitemap[req.Host]
		path := req.URL.Path

		for _, ifblock := range site.Config.Ifs {
			if ifblock.Path != nil && *ifblock.Path == path {
				for key, value := range ifblock.Respond.Headers {
					res.SetHeader(key, value)
					res.Write(500, "", nil)
					return
				}
			}
		}
	})
		
	e.Server.Use(func(res *libserve.Response, req *http.Request) {
		site := sitemap[req.Host]
		path := req.URL.Path

		if strings.HasSuffix(path, "/") {
			path = filepath.Join(path, "index.html")
		}
		path = strings.TrimPrefix(path, "/")

		f, err := site.Dist.Open(path)
		if err != nil {
			res.Write(404, "", nil)
			return
		}
		res.Write(200, path, f)
	})

	return e.Server.Serve()
}
