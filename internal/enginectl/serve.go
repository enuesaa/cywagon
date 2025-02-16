package enginectl

import (
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) Serve(confs []model.Conf) error {
	var sites []libserve.Site

	for _, conf := range confs {
		sites = append(sites, libserve.Site{
			Host:      conf.Host,
			OriginUrl: conf.Entry.Host,
			Handler:   conf.RunHandler,
		})
	}

	server := libserve.New()
	server.Port = 3000
	server.Sites = sites

	return server.Serve()
}
