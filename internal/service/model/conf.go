package model

import "github.com/zclconf/go-cty/cty"

type Config struct {
	Server Server  `hcl:"server,block"`
	Sites  []Site  `hcl:"site,block"`
	Consts []Const `hcl:"const,block"`
}

type Server struct {
	Port uint `hcl:"port,optional"`
}

type Site struct {
	Name string `hcl:"name,label"`
	Host string `hcl:"host"`
	Dist string `hcl:"dist"`
	Path []Path `hcl:"path,block"`
}

type Path struct {
	Pattern string `hcl:"pattern"`
	Status  uint   `hcl:"status"`
	Body    string `hcl:"body,optional"`
	Headers map[string]string `hcl:"headers,optional"`
}

type Const struct {
	Name  string               `hcl:"name,label"`
	Attrs map[string]cty.Value `hcl:",remain"`
}
