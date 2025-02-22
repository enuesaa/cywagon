package enginectl

import (
	"fmt"

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
			url := fmt.Sprintf("%s%s", conf.Origin.Host(), conf.HealthCheck.Path)
			fetcher.CheckHttpFetch(url)
		case "TCP":
			fetcher.CheckTcpConn(conf.Origin.Url)
		}
	}
}
