package enginectl

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libfetch"
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
	"go.uber.org/mock/gomock"
)

type EngineInterface interface {
	PrintBanner(confs []model.Conf)
	StartHealthCheck(confs []model.Conf) error
	Serve(confs []model.Conf) error
	StartUp(confs []model.Conf) error
	ValidateConfs(confs []model.Conf) error
}

func New() *Engine {
	engine := Engine{
		Container: infra.Default,
		Server: libserve.New(),
		ConfSrv: service.NewConfSrv(),
		Fetcher: libfetch.New(),
	}
	return &engine
}

type Engine struct {
	infra.Container

	Server libserve.Server
	ConfSrv service.ConfSrvInterface
	Fetcher libfetch.Fetcher
}

func NewMock(t *testing.T, prepares... func(*MockEngineInterface)) EngineInterface {
	ctrl := gomock.NewController(t)
	mock := NewMockEngineInterface(ctrl)

	for _, prepare := range prepares {
		prepare(mock)
	}
	return mock
}
