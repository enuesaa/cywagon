package enginectl

import (
	// "github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) Serve(confs []model.Config) error {
	// for _, conf := range confs {
	// 	site := libserve.Site{
	// 		// Host:      conf.Host,
	// 		// Handler: func(res *libserve.HandlerResponse, next libserve.Next, req libserve.HandlerRequest) error {
	// 		// 	// ここで in memory からデータを読み取る
				
	// 		// 	return conf.Handler(res, next, req)
	// 		// },
	// 		// Cache:     true,
	// 	}
	// 	e.Server.Sites.Push(site)
	// }

	e.Server.Port = 3000

	return e.Server.Serve()
}
