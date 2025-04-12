package handle

import (
	"log"
	"os"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
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
	filename := "./testdata/main.hcl"
	fbytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL(fbytes, filename)
	if diags.HasErrors() {
		return diags
	}

	var config Config
	confDiags := gohcl.DecodeBody(file.Body, nil, &config)

	if confDiags.HasErrors() {
		return confDiags
	}
	log.Printf("Configuration is %#v", config)

	return nil
}
