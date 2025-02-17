package service

import (
	"fmt"
	"strings"

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

	var ret []string
	for _, item := range list {
		if strings.HasSuffix(item, ".lua") {
			ret = append(ret, item)
		}
	}
	return ret
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
		Host: "",
		Entry: model.ConfEntry{
			Workdir:        ".",
			Cmd:            "",
			WaitForHealthy: 5,
		},
		HealthCheck: model.ConfHealthCheck{
			Protocol: "HTTP",
			Method:   "GET",
			Path:     "/",
		},
		Handler: nil,
	}
	runner := liblua.New()

	if err := runner.Inject(config); err != nil {
		return config, err
	}
	if err := runner.Run(code); err != nil {
		return config, err
	}
	if err := runner.Eject(&config); err != nil {
		return config, err
	}
	return config, nil
}

var ErrConfHostRequired = fmt.Errorf("host is required")

func (c *ConfService) Validate(conf model.Conf) error {
	if conf.Host == "" {
		return ErrConfHostRequired
	}
	return nil
}
