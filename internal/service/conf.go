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
	ReadHCLFiles(fpaths []string) (hcl.Body, error)
	Read(fpaths []string) (model.Config, error)
	Format(fpaths []string) error
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

func (c *ConfSrv) Read(fpaths []string) (model.Config, error) {
	var config model.Config

	hclbody, err := c.ReadHCLFiles(fpaths)
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
	c.applyDefault(&config)

	return config, nil
}

func (c *ConfSrv) applyDefault(config *model.Config) {
	if config.Server.LogFile == nil {
		def := "/dev/stdout"
		config.Server.LogFile = &def
	}
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
