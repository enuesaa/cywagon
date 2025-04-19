package enginectl

import (
	"fmt"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

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

	e.Server.Use(func(w http.ResponseWriter, req *http.Request) error {
		host := req.Host
		_, ok := sitemap[host]
		if !ok {
			w.WriteHeader(500)
			return fmt.Errorf("end")
		}
		fmt.Printf("found: %s\n", host)
		return nil
	})

	e.Server.Use(func(w http.ResponseWriter, req *http.Request) error {
		site := sitemap[req.Host]
		path := req.URL.Path

		for _, ifblock := range site.Config.Ifs {
			if ifblock.Path != nil && *ifblock.Path == path {
				for key, value := range ifblock.Respond.Headers {
					w.Header().Set(key, value)
					w.WriteHeader(200)
					return fmt.Errorf("end")
				}
			}
		}
		return nil
	})
		
	e.Server.Use(func(w http.ResponseWriter, req *http.Request) error {
			site := sitemap[req.Host]
			path := req.URL.Path

		if strings.HasSuffix(path, "/") {
			path = filepath.Join(path, "index.html")
		}
		path = strings.TrimPrefix(path, "/")

		f, err := site.Dist.Open(path)
		if err != nil {
			w.WriteHeader(404)
			return fmt.Errorf("end")
		}

		fbytes, err := io.ReadAll(f)
		if err != nil {
			w.WriteHeader(404)
			return fmt.Errorf("end")
		}

		ext := filepath.Ext(path)
		w.Header().Set("Content-Type", mime.TypeByExtension(ext))

		w.WriteHeader(200)
		w.Write(fbytes)

		return nil
	})

	return e.Server.Serve()
}
