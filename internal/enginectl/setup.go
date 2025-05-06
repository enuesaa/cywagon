package enginectl

import (
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) Setup(config model.Config, workdir string) error {	
	if err := e.Log.SetLogFile(*config.Server.LogFile); err != nil {
		return err
	}
	e.Server.Port = config.Server.Port

	e.loadSites(config)
	e.loadLogics(config)

	if err := e.loadDists(config, workdir); err != nil {
		return err
	}
	return nil
}
