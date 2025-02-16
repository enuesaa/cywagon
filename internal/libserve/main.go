package libserve

import (
	"net/url"

	"github.com/enuesaa/cywagon/internal/infra"
)

func New() Server {
	return Server{
		Container: infra.Default,
	}
}

type Server struct {
	infra.Container

	Port    int
	Sites   []Site
}

type Site struct {
	Host            string // Example: `example.com`
	OriginUrl       string // Example: `https://example.com`
	Handler         FnHandler
	parsedOriginUrl *url.URL
}
type FnHandler func(*FnHandlerResponse, FnNext, FnHandlerRequest) error
type FnNext func(FnHandlerRequest) FnHandlerResponse

type FnHandlerRequest struct {
	Path string
}
type FnHandlerResponse struct {
	Status int
}
