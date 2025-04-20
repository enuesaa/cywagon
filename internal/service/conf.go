package service

import (
	"path/filepath"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libhcl"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func NewConfSrv() ConfSrvInterface {
	return &ConfSrv{
		Container: infra.Default,
		Hcl: libhcl.New(),
	}
}

type ConfSrvInterface interface {
	ReadInWorkdir(workdir string) (model.Config, error)
}

type ConfSrv struct {
	infra.Container

	Hcl libhcl.Parser
}

func (c *ConfSrv) ReadInWorkdir(workdir string) (model.Config, error) {
	var config model.Config

	path := filepath.Join(workdir, "server.hcl")
	fbytes, err := c.Fs.Read(path)
	if err != nil {
		return config, err
	}
	if err := c.Hcl.Parse(fbytes, &config); err != nil {
		return config, err
	}
	return config, nil
}
