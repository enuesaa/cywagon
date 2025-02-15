package enginectl

import (
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) Serve(confs []model.Conf) error {
	var sites []libserve.ServeOptsSite

	for _, conf := range confs {
		sites = append(sites, libserve.ServeOptsSite{
			Host:      conf.Host,
			OriginUrl: conf.Entry.Host,
			Handler:   conf.RunHandler,
		})
	}

	serveOpts := libserve.ServeOpts{
		Port:  3000,
		Sites: sites,
	}
	return libserve.Serve(serveOpts)
}
