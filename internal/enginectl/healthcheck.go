package enginectl

import (
	"time"

	"github.com/enuesaa/cywagon/internal/libfetch"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) StartHealthCheck(confs []model.Conf) error {
	go e.poolHealthCheck(confs)
	return nil
}

func (e *Engine) poolHealthCheck(confs []model.Conf) {
	for {
		time.Sleep(1 * time.Second)
		if err := e.runHealthCheck(confs); err != nil {
			e.Log.Error(err)
		}
	}
}

func (e *Engine) runHealthCheck(confs []model.Conf) error {
	fetcher := libfetch.New()

	for _, conf := range confs {
		switch conf.HealthCheck.Protocol {
		case "HTTP":
			if err := fetcher.CheckHTTP(conf.HealthCheckUrl(), conf.HealthCheck.Matcher); err != nil {
				return err
			}
		case "TCP":
			if err := fetcher.CheckTCP(conf.Origin.Url); err != nil {
				return err
			}
		}
	}
	return nil
}
