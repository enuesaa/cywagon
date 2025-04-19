package libserve

import "github.com/enuesaa/cywagon/internal/infra"

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
