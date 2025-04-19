package libserve

type Handler func(c *Context) *Response

func (s *Server) Use(handler Handler) {
	s.handlers = append(s.handlers, handler)
}
