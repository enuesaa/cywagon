package libserve

import (
	"fmt"
	"io/fs"

	"github.com/enuesaa/cywagon/internal/infra"
)

func New() Server {
	return Server{
		Container: infra.Default,
		Port: 3000,
		Sites: map[string]Site{},
	}
}

type Server struct {
	infra.Container

	Port  int
	Sites Sites
}

type Site struct {
	Host    string // Example: `example.com`
	Handler Handler
	Dist    fs.FS
}
type Handler func(*HandlerResponse, Next, HandlerRequest) error
type Next func(HandlerRequest) HandlerResponse

type HandlerRequest struct {
	Path string
}
type HandlerResponse struct {
	Status int
}

type Sites map[string]Site

func (m *Sites) Push(site Site) {
	if len(*m) == 0 {
		(*m)["default"] = site
	}
	(*m)[site.Host] = site
}

func (m *Sites) getByHost(host string) Site {
	site, ok := (*m)[host]
	if ok {
		return site
	}
	return (*m)["default"]
}

var ErrSitesNeedAtLeast1SiteDef = fmt.Errorf("sites need at least 1 def")

func (m *Sites) Validate() error {
	if _, ok := (*m)["default"]; !ok {
		return ErrSitesNeedAtLeast1SiteDef
	}
	return nil
}
