package enginectl

import (
	"fmt"
	"time"

	"github.com/enuesaa/cywagon/internal/libfetch"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) StartHealthCheck(confs []model.Conf) error {
	for _, conf := range confs {
		if conf.HealthCheck.Protocol == "HTTP" {
			go e.checkHealthByHTTP(conf)
		}
		switch conf.HealthCheck.Protocol {
		case "HTTP":
			go e.checkHealthByHTTP(conf)
		case "TCP":
			go e.checkHealthByTCP(conf)
		}
	}
	return nil
}

func (e *Engine) checkHealthByHTTP(conf model.Conf) {
	fetcher := libfetch.New()

	url := fmt.Sprintf("%s%s", conf.Entry.Host, conf.HealthCheck.Path)

	for {
		time.Sleep(time.Duration(conf.Entry.WaitForHealthy) * time.Second)
		status := fetcher.CheckHttpFetch(url)

		e.Log.Info("[healthcheck] %d", status)
	}
}

func (e *Engine) checkHealthByTCP(conf model.Conf) {
	fetcher := libfetch.New()

	for {
		time.Sleep(time.Duration(conf.Entry.WaitForHealthy) * time.Second)
		status := fetcher.CheckTcpConn(conf.Entry.Host)

		e.Log.Info("[healthcheck] %d", status)
	}
}
