package model

import "github.com/enuesaa/cywagon/internal/liblua"

// deprecated
type Conf struct {
	Host        string          `lua:"host"`
	Handler     liblua.Fn       `lua:"handler"`
}

type Config struct {
	Server Server `hcl:"server,block"`
	Sites  []Site `hcl:"site,block"`
}

type Server struct {
	Port uint `hcl:"port"`
}

type Site struct {
	Name string `hcl:"name,label"`
	Host string `hcl:"host"`
	Dist string `hcl:"dist"`
}
