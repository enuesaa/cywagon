package service

import (
	"strings"

	"github.com/creasty/defaults"
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
	ListHCLFiles(dir string) ([]string, error)
	ReadHCLFiles(fpaths []string) (hcl.Body, error)
	Parse(hclbody hcl.Body) (model.Config, error)
	ReadInDir(dir string) (model.Config, error)
	Format(fpaths []string) error
}

type ConfSrv struct {
	infra.Container

	Hcl libhcl.Parser
}

func (c *ConfSrv) ListHCLFiles(dir string) ([]string, error) {
	fpaths, err := c.Fs.ListFiles(dir)
	if err != nil {
		return nil, err
	}
	var list []string

	for _, fpath := range fpaths {
		if strings.HasSuffix(fpath, ".hcl") {
			list = append(list, fpath)
		}
	}
	return list, nil
}

func (c *ConfSrv) ReadHCLFiles(fpaths []string) (hcl.Body, error) {
	files := make(map[string][]byte, 0)

	for _, fpath := range fpaths {
		fbytes, err := c.Fs.Read(fpath)
		if err != nil {
			return nil, err
		}
		files[fpath] = fbytes
	}
	return c.Hcl.MergeHCLFiles(files)
}

func (c *ConfSrv) Parse(hclbody hcl.Body) (model.Config, error) {
	var config model.Config

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
	if err := c.applyDefaults(&config); err != nil {
		return config, err
	}
	return config, nil
}

func (c *ConfSrv) applyDefaults(config *model.Config) error {
	return defaults.Set(config)
}

func (c *ConfSrv) ReadInDir(dir string) (model.Config, error) {
	fpaths, err := c.ListHCLFiles(dir)
	if err != nil {
		return model.Config{}, err
	}
	hclbody, err := c.ReadHCLFiles(fpaths)
	if err != nil {
		return model.Config{}, err
	}
	return c.Parse(hclbody)
}

func (c *ConfSrv) Format(fpaths []string) error {
	for _, fpath := range fpaths {
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
