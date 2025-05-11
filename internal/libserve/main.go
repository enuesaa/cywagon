package libserve

import "github.com/enuesaa/cywagon/internal/infra"

func New() Server {
	return Server{
		Container:  infra.Default,
		OnResponse: func(c *Context, status int, method string) {},
		OnError:    func(c *Context, err error) {},
		handlers:   make([]Handler, 0),
	}
}

type Server struct {
	infra.Container

	handlers   []Handler
	listenmap  map[int][]ListenConfig
	OnResponse FnOnResponse
	OnError    FnOnError
}

type Handler func(c *Context) *Response
type FnOnResponse func(c *Context, status int, method string)
type FnOnError func(c *Context, err error)

type ListenConfig struct {
	tls bool
	certfile string
	keyfile string
}
