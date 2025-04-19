package enginectl

import (
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) Serve(config model.Config) error {
	e.Server.Port = config.Server.Port

	for _, site := range config.Sites {
		e.LoadFS(site.Host, site.Dist)

		ssite := libserve.Site{
			Host: site.Host,
			Dist: e.dists[site.Host],
		}
		e.Server.Push(ssite)
	}

	return e.Server.Serve()
}
