package service

import (
	"maps"
	"strings"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libhcl"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
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

func (c *ConfSrv) ListHCLFiles(workdir string) (map[string][]byte, error) {
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
	return files, nil
}

func (c *ConfSrv) Read(workdir string) (model.Config, error) {
	var config model.Config

	files, err := c.ListHCLFiles(workdir)
	if err != nil {
		return config, err
	}
	hclbody, err := c.Hcl.MergeHCLFiles(files)
	if err != nil {
		return config, err
	}

	type PartialConstsConfig struct {
		Consts []model.Const `hcl:"const,block"`
		Remain hcl.Body      `hcl:",remain"`
	}
	var partialconsts PartialConstsConfig

	if err := c.Hcl.Decode(hclbody, nil, &partialconsts); err != nil {
		return config, err
	}
	constsmap := make(map[string]cty.Value)
	for _, co := range partialconsts.Consts {
		maps.Copy(constsmap, co.Attrs)
	}

	vars := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"const": cty.ObjectVal(constsmap),
		},
	}
	if err := c.Hcl.Decode(hclbody, vars, &config); err != nil {
		return config, err
	}
	return config, nil
}
