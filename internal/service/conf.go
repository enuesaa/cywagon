package service

import (
	"fmt"
	"strings"
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/liblua"
	"github.com/enuesaa/cywagon/internal/service/model"
	"go.uber.org/mock/gomock"
)

func NewConfService() ConfServicer {
	return &ConfService{
		Container: infra.Default,
	}
}

func NewConfServiceMock(t *testing.T, prepares... func(*MockConfServicer)) ConfServicer {
	ctrl := gomock.NewController(t)
	mock := NewMockConfServicer(ctrl)
	for _, prepare := range prepares {
		prepare(mock)
	}
	return mock
}

type ConfServicer interface {
	List(search []string) ([]model.Conf, error)
	ListConfPaths(search []string) ([]string, error)
	IsConfPath(path string) bool
	Read(path string) (model.Conf, error)
	Validate(conf model.Conf) error
}

type ConfService struct {
	infra.Container
}

func (c *ConfService) List(search []string) ([]model.Conf, error) {
	confpaths, err := c.ListConfPaths(search)
	if err != nil {
		return nil, err
	}
	var list []model.Conf

	for _, confpath := range confpaths {
		conf, err := c.Read(confpath)
		if err != nil {
			return nil, err
		}
		list = append(list, conf)
	}
	return list, nil
}

func (c *ConfService) ListConfPaths(search []string) ([]string, error) {
	var list []string

	for _, path := range search {
		if !c.Fs.IsExist(path) {
			return nil, fmt.Errorf("path not found: %s", path)
		}
		if c.Fs.IsFile(path) {
			list = append(list, path)
			continue
		}
		files, err := c.Fs.ListFiles(path)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			if c.IsConfPath(file) {
				list = append(list, file)
			}
		}
	}
	return list, nil
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
