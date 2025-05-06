package libserve

import (
	"fmt"
	"net/http"
)

type Handler func(c *Context) *Response

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
			// log
			// s.Log.Info("%d %s %s", res.status, ctx.req.Method, ctx.req.URL.Path)

			// flush
			if err := res.flush(w); err != nil {
				// s.Log.Info("Error: %w", err)
			}
			break
		}
	}
}
