package enginectl

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/libsock"
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
	"go.uber.org/mock/gomock"
)

type EngineInterface interface {
	PrintBanner(confs []model.Conf)
	Serve(confs []model.Conf) error
	ValidateConfs(confs []model.Conf) error
	Deploy(sitename string, path string) error
	Read(sitename string) (string, error)
	StartListenSock() error
	Close() error
}

func New() *Engine {
	engine := Engine{
		Container: infra.Default,
		Server: libserve.New(),
		ConfSrv: service.NewConfSrv(),
		Sock: libsock.New(),
		dists: make(Dists),
	}
	return &engine
}

type Engine struct {
	infra.Container

	Server libserve.Server
	ConfSrv service.ConfSrvInterface
	Sock libsock.Sock
	dists Dists
}

func NewMock(t *testing.T, prepares... func(*MockEngineInterface)) EngineInterface {
	ctrl := gomock.NewController(t)
	mock := NewMockEngineInterface(ctrl)

	for _, prepare := range prepares {
		prepare(mock)
	}
	return mock
}
