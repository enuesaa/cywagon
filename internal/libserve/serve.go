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
	ctx := NewContext(req)

	for _, handler := range s.handlers {
		res := handler(&ctx)
		if res != nil {
			if err := res.flush(w); err != nil {
				s.Log.Info("Error: %w", err)
			}
			break
		}
	}
}
