package libserve

import (
	"fmt"
	"net/http"
)

type Handler func(c *Context) *Response
type Logger func(c *Context, res *Response)

func (s *Server) Use(handler Handler) {
	s.handlers = append(s.handlers, handler)
}

func (s *Server) UseLogger(logger Logger) {
	s.logger = logger
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
			s.logger(&ctx, res)

			// flush
			if err := res.flush(w); err != nil {
				// s.Log.Info("Error: %w", err)
			}
			break
		}
	}
}
