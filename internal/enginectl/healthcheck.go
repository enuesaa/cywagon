package enginectl

import (
	"time"

	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) StartHealthCheck(confs []model.Conf) error {
	wait := e.calcMaxWaitForHealthy(confs)
	time.Sleep(time.Duration(wait) * time.Second)

	// first check must healthy
	if err := e.runHealthCheck(confs); err != nil {
		return err
	}
	go e.poolHealthCheck(confs)

	return nil
}

func (e *Engine) calcMaxWaitForHealthy(confs []model.Conf) int {
	wait := 1
	for _, conf := range confs {
		if conf.Origin.WaitForHealthy > wait {
			wait = conf.Origin.WaitForHealthy
		}
	}
	return wait
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
	for _, conf := range confs {
		switch conf.HealthCheck.Protocol {
		case "HTTP":
			if err := e.Fetcher.CheckHTTP(conf.HealthCheckUrl(), conf.HealthCheck.Matcher); err != nil {
				return err
			}
		case "TCP":
			if err := e.Fetcher.CheckTCP(conf.Origin.Url); err != nil {
				return err
			}
		}
	}
	return nil
}
