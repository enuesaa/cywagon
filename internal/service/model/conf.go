package model

import "github.com/zclconf/go-cty/cty"

type Config struct {
	Server Server  `hcl:"server,block"`
	Sites  []Site  `hcl:"site,block"`
	Consts []Const `hcl:"const,block"`
}

type Server struct {
	Port int `hcl:"port"`
}

type Const struct {
	Attrs map[string]cty.Value `hcl:",remain"`
}

type Site struct {
	Name    string            `hcl:"name,label"`
	Host    string            `hcl:"host"`
	Dist    string            `hcl:"dist"`
	Headers map[string]string `hcl:"headers,optional"`
	Ifs     []If              `hcl:"if,block"`
}

type If struct {
	Path      *string   `hcl:"path,optional"`
	PathIn    []string `hcl:"path_in,optional"`
	PathNot   *string   `hcl:"path_not,optional"`
	PathNotIn []string `hcl:"path_not_in,optional"`

	Headers      map[string]string   `hcl:"headers,optional"`
	HeadersIn    []map[string]string `hcl:"headers_in,optional"`
	HeadersNot   map[string]string   `hcl:"headers_not,optional"`
	HeadersNotIn []map[string]string `hcl:"headers_not_in,optional"`

	Ipaddr      *string   `hcl:"ipaddr,optional"`
	IpaddrNot   *string   `hcl:"ipaddr_not,optional"`
	IpaddrIn    []string `hcl:"ipaddr_in,optional"`
	IpaddrNotIn []string `hcl:"ipaddr_not_in,optional"`

	Respond Respond `hcl:"respond,block"`
}

type Respond struct {
	Status  *int              `hcl:"status,optional"`
	Headers map[string]string `hcl:"headers,optional"`
	Body    *string           `hcl:"body,optional"`
}
