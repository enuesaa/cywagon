package libserve

import (
	"fmt"
	"net/http"
)

func (s *Server) Serve() error {
	if err := s.Validate(); err != nil {
		return err
	}
	addr := fmt.Sprintf(":%d", s.port)

	return http.ListenAndServe(addr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	host := req.Host

	site := s.getByHost(host)
	http.ServeFileFS(w, req, site.Dist, ".")
}
