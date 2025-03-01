package libserve

import "github.com/enuesaa/cywagon/internal/infra"

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
