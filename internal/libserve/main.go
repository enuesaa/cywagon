package libserve

import (
	"fmt"

	"github.com/enuesaa/cywagon/internal/infra"
)

func New() Server {
	return Server{
		Container: infra.Default,
		Port:      3000,
		sites:     map[string]Site{},
	}
}

type Server struct {
	infra.Container

	Port  int
	sites map[string]Site
}

func (s *Server) Add(site Site) {
	if len(s.sites) == 0 {
		s.sites["default"] = site
	}
	s.sites[site.Host] = site
}

func (s *Server) getByHost(host string) Site {
	site, ok := s.sites[host]
	if ok {
		return site
	}
	return s.sites["default"]
}

func (s *Server) Validate() error {
	if _, ok := s.sites["default"]; !ok {
		return fmt.Errorf("sites need at least 1 def")
	}
	return nil
}
