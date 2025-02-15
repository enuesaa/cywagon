package enginectl

import (
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func Serve(ctn infra.Container, confs []model.Conf) error {
	sites := make([]libserve.ServeOptsSite, 0)
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
