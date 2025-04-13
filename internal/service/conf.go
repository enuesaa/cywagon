package service

import (
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libhcl"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func NewConfSrv() ConfSrvInterface {
	return &ConfSrv{
		Container: infra.Default,
	}
}

type ConfSrvInterface interface {
	Read(path string) (model.Config, error)
}

type ConfSrv struct {
	infra.Container

	Hcl libhcl.Parser
}

func (c *ConfSrv) Read(path string) (model.Config, error) {
	var config model.Config

	fbytes, err := c.Fs.Read(path)
	if err != nil {
		return config, err
	}
	if err := c.Hcl.Parse(fbytes, &config); err != nil {
		return config, err
	}
	c.applyDefault(&config)

	return config, nil
}

func (c *ConfSrv) applyDefault(config *model.Config) {
	if config.Server.Port == 0 {
		config.Server.Port = 3000
	}
}
