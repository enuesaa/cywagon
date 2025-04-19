package libserve

import (
	"fmt"
	"net/http"
)

func (s *Server) Serve() error {
	addr := fmt.Sprintf(":%d", s.Port)

	return http.ListenAndServe(addr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	res := NewResponse()

	for _, handler := range s.handlers {
		handler(&res, req)
		if res.Close() {
			break
		}
	}

	if err := res.Flush(w); err != nil {
		s.Log.Info("Error: %w", err)
	}
}
