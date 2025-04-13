package service

import (
	"fmt"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libhcl"
	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
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

	defMap := make(map[string]cty.Value)
	for _, def := range config.Defs {
		defMap[def.Name] = cty.ObjectVal(def.Props)
	}

	a := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"def": cty.ObjectVal(defMap),
		},
	}
	fmt.Printf("%+v\n", a)
	

	return config, nil
}

func (c *ConfSrv) applyDefault(config *model.Config) {
	if config.Server.Port == 0 {
		config.Server.Port = 3000
	}
}
