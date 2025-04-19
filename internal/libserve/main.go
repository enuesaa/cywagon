package libserve

import (
	"net/http"

	"github.com/enuesaa/cywagon/internal/infra"
)

func New() Server {
	return Server{
		Container: infra.Default,
		Port:      3000,
		handlers:  make([]Handler, 0),
	}
}

type Server struct {
	infra.Container

	Port  int
	handlers []Handler
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (s *Server) Use(handler Handler) {
	s.handlers = append(s.handlers, handler)
}
