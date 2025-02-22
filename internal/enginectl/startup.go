package enginectl

import "github.com/enuesaa/cywagon/internal/service/model"

func (e *Engine) StartUp(confs []model.Conf) error {
	for _, conf := range confs {
		if conf.Entry.Cmd != "" {
			go e.runCmd(conf.Entry.Workdir, conf.Entry.Cmd)
		}
		// TODO: unregister other sites here.
	}
	return nil
}

func (e *Engine) runCmd(workdir string, command string) {
	if err := e.Cmd.Start(workdir, command); err != nil {
		e.Log.Error(err)
	}
}
