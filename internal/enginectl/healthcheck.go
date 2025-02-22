package enginectl

import (
	"github.com/enuesaa/cywagon/internal/libfetch"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) StartHealthCheck(confs []model.Conf) error {
	go e.runHealthCheck(confs)
	return nil
}

func (e *Engine) runHealthCheck(confs []model.Conf) {
	fetcher := libfetch.New()

	for _, conf := range confs {
		switch conf.HealthCheck.Protocol {
		case "HTTP":
			fetcher.CheckHttpFetch(conf.HealthCheck.Protocol, conf.Entry.Host, conf.HealthCheck.Path)
		case "TCP":
			fetcher.CheckTcpConn(conf.Entry.Host)
		}
	}
}
