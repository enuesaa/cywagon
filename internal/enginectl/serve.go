package enginectl

import (
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) Serve(confs []model.Conf) error {
	for _, conf := range confs {
		site := libserve.Site{
			Host:      conf.Host,
			OriginUrl: conf.Origin.Url,
			Handler: func(res *libserve.HandlerResponse, next libserve.Next, req libserve.HandlerRequest) error {
				return conf.Handler(res, next, req)
			},
			Cache:     conf.Cache,
		}
		if err := e.Server.Sites.Push(site); err != nil {
			return err
		}
	}

	e.Server.Port = 3000

	return e.Server.Serve()
}
