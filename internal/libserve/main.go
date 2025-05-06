package libserve

import "github.com/enuesaa/cywagon/internal/infra"

func New() Server {
	return Server{
		Container: infra.Default,
		Port:      3000,
		handlers:  make([]Handler, 0),
		logger:    func(c *Context, res *Response) {},
	}
}

type Server struct {
	infra.Container

	Port     int
	handlers []Handler
	logger   Logger
}
