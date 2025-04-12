package handle

import (
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

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

func Check() error {
	var config Config
	if err := hclsimple.DecodeFile("testdata/main.hcl", nil, &config); err != nil {
		return err
	}
	log.Printf("Configuration is %#v", config)

	return nil
}