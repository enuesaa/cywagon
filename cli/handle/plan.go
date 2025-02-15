package handle

import "github.com/enuesaa/cywagon/internal/enginectl"

func Plan(confDir string) error {
	engine := enginectl.New()

	return engine.Validate(confDir)
}
