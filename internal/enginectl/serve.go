package enginectl

import "github.com/enuesaa/cywagon/internal/service/model"

func (e *Engine) Serve(config model.Config) error {
	e.Server.Port = int(config.Server.Port)

	return e.Server.Serve()
}
