package enginectl

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) Serve(config model.Config) error {
	e.Server.Port = config.Server.Port

	for _, site := range config.Sites {
		dist, err := e.LoadFS(site.Host, site.Dist)
		if err != nil {
			return err
		}

		ssite := libserve.Site{
			Host: site.Host,
			Handle: func(w http.ResponseWriter, req *http.Request) {
				path := req.URL.Path
				fmt.Println(path)

				if strings.HasSuffix(path, "/") {
					path = filepath.Join(path, "index.html")
				}
				path = strings.TrimPrefix(path, "/")

				f, err := dist.Open(path)
				if err != nil {
					w.WriteHeader(404)
					return
				}

				fbytes, err := io.ReadAll(f)
				if err != nil {
					w.WriteHeader(404)
					return
				}

				// ext
				ext := filepath.Ext(path)
				w.Header().Set("Content-Type", mime.TypeByExtension(ext))

				w.WriteHeader(200)
				w.Write(fbytes)
			},
		}
		e.Server.Add(ssite)
	}

	return e.Server.Serve()
}
