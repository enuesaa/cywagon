package libfetch

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"go.uber.org/mock/gomock"
)

type FetcherInterface interface {
	FetchHTTP(url string) string
	ConnectTCP(address string) error
	CheckHTTP(url string, matcher string) error
	CheckTCP(address string) error
}

func New() Fetcher {
	return Fetcher{
		Container: infra.Default,
	}
}

type Fetcher struct {
	infra.Container
}

func NewMock(t *testing.T, prepares... func(*MockFetcherInterface)) FetcherInterface {
	ctrl := gomock.NewController(t)
	mock := NewMockFetcherInterface(ctrl)

	for _, prepare := range prepares {
		prepare(mock)
	}
	return mock
}
