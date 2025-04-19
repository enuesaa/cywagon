package libserve

import "net/http"

type Handler func(res *Response, req *http.Request)

func (s *Server) Use(handler Handler) {
	s.handlers = append(s.handlers, handler)
}
