package handle

import (
	"github.com/enuesaa/cywagon/internal/enginectl"
	"github.com/enuesaa/cywagon/internal/infra"
)

func Plan(ctn infra.Container, confDir string) error {
	engine := enginectl.New(ctn)
	
	return engine.Validate(confDir)
}
