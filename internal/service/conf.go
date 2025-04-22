package service

import (
	"strings"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libhcl"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/hashicorp/hcl/v2"
)

func NewConfSrv() ConfSrvInterface {
	return &ConfSrv{
		Container: infra.Default,
		Hcl: libhcl.New(),
	}
}

type ConfSrvInterface interface {
	Read(workdir string) (model.Config, error)
}

type ConfSrv struct {
	infra.Container

	Hcl libhcl.Parser
}

func (c *ConfSrv) ReadHCLFiles(workdir string) (hcl.Body, error) {
	files := make(map[string][]byte, 0)
	fpaths, err := c.Fs.ListFiles(workdir)
	if err != nil {
		return nil, err
	}
	for _, fpath := range fpaths {
		if !strings.HasSuffix(fpath, ".hcl") {
			continue
		}
		fbytes, err := c.Fs.Read(fpath)
		if err != nil {
			return nil, err
		}
		files[fpath] = fbytes
	}
	return c.Hcl.MergeHCLFiles(files)
}

func (c *ConfSrv) Read(workdir string) (model.Config, error) {
	var config model.Config

	hclbody, err := c.ReadHCLFiles(workdir)
	if err != nil {
		return config, err
	}
	var partialconsts model.PartialConstsConfig

	if err := c.Hcl.Decode(hclbody, &partialconsts); err != nil {
		return config, err
	}
	c.Hcl.UseVar("const", partialconsts.FlattenConsts())

	if err := c.Hcl.Decode(hclbody, &config); err != nil {
		return config, err
	}
	return config, nil
}
