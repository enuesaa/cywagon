package service

import (
	"fmt"
	"strings"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/liblua"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func NewConfService(container infra.Container) ConfService {
	return ConfService{
		Container: container,
	}
}

type ConfService struct {
	infra.Container
}

func (c *ConfService) IsConfPath(path string) bool {
	return strings.HasSuffix(path, ".lua")
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
		Origin: model.ConfOrigin{
			Workdir:        ".",
			Cmd:            "",
			WaitForHealthy: 5,
			Url:            "",
		},
		HealthCheck: model.ConfHealthCheck{
			Protocol: "HTTP",
			Method:   "GET",
			Path:     "/",
			Matcher:  "200",
		},
		Handler: nil,
	}
	runner := liblua.New(c.Container)

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
