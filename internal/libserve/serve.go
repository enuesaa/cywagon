package libserve

import (
	"fmt"
	"net/http"
)

type Handler func(c *Context) *Response
type FnOnResponse func(c *Context, status int, method string)
type FnOnError func(c *Context, err error)

func (s *Server) Use(handler Handler) {
	s.handlers = append(s.handlers, handler)
}

func (s *Server) Serve() error {
	addr := fmt.Sprintf(":%d", s.Port)

	return http.ListenAndServe(addr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := NewContext(req)

	for _, handler := range s.handlers {
		res := handler(&ctx)
		if res != nil {
			s.OnResponse(&ctx, res.status, ctx.req.Method)

			if err := res.flush(w); err != nil {
				s.OnError(&ctx, err)
			}
			break
		}
	}
}
