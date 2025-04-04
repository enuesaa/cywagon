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

func NewConfSrv() ConfSrvInterface {
	return &ConfSrv{
		Container: infra.Default,
	}
}

func NewConfSrvMock(t *testing.T, prepares... func(*MockConfSrvInterface)) *MockConfSrvInterface {
	ctrl := gomock.NewController(t)
	mock := NewMockConfSrvInterface(ctrl)
	for _, prepare := range prepares {
		prepare(mock)
	}
	return mock
}

type ConfSrvInterface interface {
	List(search []string) ([]model.Conf, error)
	ListConfPaths(search []string) ([]string, error)
	IsConfPath(path string) bool
	Read(path string) (model.Conf, error)
	Validate(conf model.Conf) error
}

type ConfSrv struct {
	infra.Container
}

func (c *ConfSrv) List(search []string) ([]model.Conf, error) {
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

func (c *ConfSrv) ListConfPaths(search []string) ([]string, error) {
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

func (c *ConfSrv) IsConfPath(path string) bool {
	return strings.HasSuffix(path, ".lua")
}

func (c *ConfSrv) Read(path string) (model.Conf, error) {
	codeb, err := c.Fs.Read(path)
	if err != nil {
		return model.Conf{}, err
	}
	code := string(codeb)

	return c.parse(code)
}

func (c *ConfSrv) parse(code string) (model.Conf, error) {
	config := model.Conf{
		Host: "",
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

func (c *ConfSrv) Validate(conf model.Conf) error {
	if conf.Host == "" {
		return ErrConfHostRequired
	}
	return nil
}
