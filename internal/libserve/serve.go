package libserve

import (
	"fmt"
	"net/http"
)

func (s *Server) Serve() error {
	addr := fmt.Sprintf(":%d", s.Port)

	handler := Handler{
		sites: s.sites,
	}
	return http.ListenAndServe(addr, &handler)
}
