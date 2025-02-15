package service

import (
	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/liblua"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func NewConfService() ConfService {
	return ConfService{
		Container: infra.Default,
	}
}

type ConfService struct {
	infra.Container
}

func (c *ConfService) List(dir string) []string {
	list, err := c.Fs.ListFiles(dir)
	if err != nil {
		return []string{}
	}
	return list
}

func (c *ConfService) Read(path string) (model.Conf, error) {
	codeb, err := c.Fs.Read(path)
	if err != nil {
		return model.Conf{}, err
	}
	code := string(codeb)

	return c.parse(code)
}

func (c *ConfService) parse(code string) (model.Conf, error) {
	config := model.Conf{
		Host: "aa",
		Entry: model.ConfEntry{
			Workdir:        ".",
			Cmd:            "",
			WaitForHealthy: 60,
		},
		HealthCheck: model.ConfHealthCheck{
			Protocol: "HTTP",
			Method:   "GET",
			Path:     "/",
		},
		Handler: nil,
	}
	runner := liblua.NewRunner(code)

	if err := runner.Inject(config); err != nil {
		return config, err
	}
	if err := runner.Run(); err != nil {
		return config, err
	}
	if err := runner.Eject(&config); err != nil {
		return config, err
	}
	return config, nil
}
