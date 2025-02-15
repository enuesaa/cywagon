package handle

import "github.com/enuesaa/cywagon/internal/enginectl"

func Check(confDir string) error {
	engine := enginectl.New()

	confs, err := engine.ListConfs(confDir)
	if err != nil {
		return err
	}

	return engine.Validate(confs)
}
