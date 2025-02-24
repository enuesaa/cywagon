package libserve

import (
	"net/url"

	"github.com/enuesaa/cywagon/internal/infra"
)

func New() Server {
	return Server{
		Container: infra.Default,
		Port: 3000,
		Sites: newSites(),
	}
}

type Server struct {
	infra.Container

	Port  int
	Sites Sites
}

type Site struct {
	Host            string // Example: `example.com`
	OriginUrl       string // Example: `https://example.com`
	Handler         Handler
	parsedOriginUrl *url.URL
}
type Handler func(*HandlerResponse, Next, HandlerRequest) error
type Next func(HandlerRequest) HandlerResponse

type HandlerRequest struct {
	Path string
}
type HandlerResponse struct {
	Status int
}
