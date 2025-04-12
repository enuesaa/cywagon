package enginectl

import (
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) Serve(config model.Config) error {
	for _, site := range config.Sites {
		e.LoadFS(site.Host, site.Dist)

		ssite := libserve.Site{
			Host: site.Host,
			// Handler: func(res *libserve.HandlerResponse, next libserve.Next, req libserve.HandlerRequest) error {
			// ここで in memory からデータを読み取る	
			// 	return conf.Handler(res, next, req)
			// },
			Dist: e.dists[site.Host],
		}
		e.Server.Sites.Push(ssite)
	}

	e.Server.Port = int(config.Server.Port)

	return e.Server.Serve()
}
