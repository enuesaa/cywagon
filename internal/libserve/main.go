package libserve

import "github.com/enuesaa/cywagon/internal/infra"

func New() Server {
	return Server{
		Container:  infra.Default,
		Port:       3000,
		OnResponse: func(c *Context, status int, method string) {},
		OnError:    func(c *Context, err error) {},
		handlers:   make([]Handler, 0),
	}
}

type Server struct {
	infra.Container

	Port       int
	handlers   []Handler
	listeners  []Listener
	OnResponse FnOnResponse
	OnError    FnOnError
}

type Handler func(c *Context) *Response
type FnOnResponse func(c *Context, status int, method string)
type FnOnError func(c *Context, err error)
