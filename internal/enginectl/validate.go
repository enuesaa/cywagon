package enginectl

import "github.com/enuesaa/cywagon/internal/service"

// TODO: validate
func (e *Engine) Validate(confDir string) error {
	confsrv := service.NewConfService(e.ctn)

	files := confsrv.List(confDir)

	for _, file := range files {
		config, err := confsrv.Read(file)
		if err != nil {
			return err
		}
		e.ctn.Log.Info("hostname: %s", config.Host)
	}
	return nil
}
