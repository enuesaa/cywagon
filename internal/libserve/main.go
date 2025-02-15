package libserve

import "github.com/enuesaa/cywagon/internal/infra"

func New() Server {
	return Server{
		Container: infra.Default,
	}
}

type Server struct {
	infra.Container
}
