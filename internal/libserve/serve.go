package libserve

import (
	"crypto/tls"
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

	// see https://gist.github.com/denji/12b3a568f092ab951456
	cert, err := tls.LoadX509KeyPair("localhost.pem", "localhost-key.pem")
	if err != nil {
		return err
	}
	tlsconfig := tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	srv := &http.Server{
		Addr:      addr,
		Handler:   s,
		TLSConfig: &tlsconfig,
	}
	return srv.ListenAndServeTLS("", "")
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
