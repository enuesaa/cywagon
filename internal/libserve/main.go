package libserve

import (
	"net/http"
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
type FnHandler func(*http.Request, FnNext) (*http.Response, error)
type FnNext func(*http.Request) *http.Response
