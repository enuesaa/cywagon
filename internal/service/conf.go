package service

import (
	"strings"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libhcl"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func NewConfSrv() ConfSrvInterface {
	return &ConfSrv{
		Container: infra.Default,
		Hcl:       libhcl.New(),
	}
}

type ConfSrvInterface interface {
	ListHCLFiles(workdir string) ([]string, error)
	Read(workdir string) (model.Config, error)
	Format(workdir string) error
}

type ConfSrv struct {
	infra.Container

	Hcl libhcl.Parser
}

func (c *ConfSrv) ListHCLFiles(workdir string) ([]string, error) {
	fpaths, err := c.Fs.ListFiles(workdir)
	if err != nil {
		return nil, err
	}
	var list []string

	for _, fpath := range fpaths {
		if !strings.HasSuffix(fpath, ".hcl") {
			continue
		}
		list = append(list, fpath)
	}
	return list, nil
}

func (c *ConfSrv) readHCLFiles(workdir string) (hcl.Body, error) {
	list, err := c.ListHCLFiles(workdir)
	if err != nil {
		return nil, err
	}
	files := make(map[string][]byte, 0)

	for _, fpath := range list {
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

	hclbody, err := c.readHCLFiles(workdir)
	if err != nil {
		return config, err
	}

	var partialConsts model.PartialConstsConfig
	if err := c.Hcl.Decode(hclbody, &partialConsts); err != nil {
		return config, err
	}
	c.Hcl.UseVar("const", partialConsts.FlattenConsts())

	var partialLogicNameOnly model.PartialLogicNameOnlyConfig
	if err := c.Hcl.Decode(hclbody, &partialLogicNameOnly); err != nil {
		return config, err
	}
	c.Hcl.UseVar("logic", partialLogicNameOnly.FlattenLogicNames())

	if err := c.Hcl.Decode(hclbody, &config); err != nil {
		return config, err
	}
	return config, nil
}

func (c *ConfSrv) Format(workdir string) error {
	list, err := c.ListHCLFiles(workdir)
	if err != nil {
		return err
	}
	for _, fpath := range list {
		fbytes, err := c.Fs.Read(fpath)
		if err != nil {
			return err
		}
		formatted := hclwrite.Format(fbytes)

		if err := c.Fs.Create(fpath, formatted); err != nil {
			return err
		}
	}
	return nil
}
