package enginectl

import (
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) Serve(config model.Config) error {
	for _, site := range config.Sites {
		site := libserve.Site{
			Host: site.Host,
			// Handler: func(res *libserve.HandlerResponse, next libserve.Next, req libserve.HandlerRequest) error {
			// 	// ここで in memory からデータを読み取る
				
			// 	return conf.Handler(res, next, req)
			// },
		}
		e.Server.Sites.Push(site)
	}

	e.Server.Port = int(config.Server.Port)

	return e.Server.Serve()
}
