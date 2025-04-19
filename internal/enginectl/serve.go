package enginectl

import (
	"net/http"

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
				http.ServeFileFS(w, req, dist, ".")
			},
		}
		e.Server.Add(ssite)
	}

	return e.Server.Serve()
}
