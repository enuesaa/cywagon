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

type EngineCtl interface {
	PrintBanner(confs []model.Conf)
	StartHealthCheck(confs []model.Conf) error
	Serve(confs []model.Conf) error
	StartUp(confs []model.Conf) error
	ValidateConfs(confs []model.Conf) error
}

func New(container infra.Container) *Engine {
	engine := Engine{
		Container: container,
		Server: libserve.New(container),
		ConfSrv: service.NewConfService(container),
		Fetcher: libfetch.New(container),
	}
	return &engine
}

type Engine struct {
	infra.Container

	Server libserve.Server
	ConfSrv service.ConfService
	Fetcher libfetch.Fetcher
}

func NewMock(t *testing.T, prepares... func(*MockEngineCtl)) EngineCtl {
	ctrl := gomock.NewController(t)

	mock := NewMockEngineCtl(ctrl)

	for _, prepare := range prepares {
		prepare(mock)
	}
	return mock
}
