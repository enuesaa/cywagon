package handle

import (
	"log"
	"os"

	"github.com/enuesaa/cywagon/internal/libhcl"
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

func (h *Handler) Check(path string) error {
	fbytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var config Config
	parser := libhcl.New()
	if err := parser.Parse(fbytes, &config); err != nil {
		return err
	}
	log.Printf("Configuration is %#v", config)

	return nil
}
