package service

import (
	"fmt"

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
	Validate(config model.Config) error
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
	return config, nil
}

var ErrConfHostRequired = fmt.Errorf("host is required")

func (c *ConfSrv) Validate(config model.Config) error {
	if config.Server.Port == 0 {
		return ErrConfHostRequired
	}
	return nil
}
