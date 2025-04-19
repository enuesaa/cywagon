package libserve

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrFlushResponse = fmt.Errorf("error")

func (s *Server) Serve() error {
	addr := fmt.Sprintf(":%d", s.Port)

	return http.ListenAndServe(addr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, handler := range s.handlers {
		if err := handler(w, req); err != nil {
			if errors.Is(err, ErrFlushResponse) {
				break
			}
			fmt.Printf("err: %s", err.Error())
			break
		}
	}
}
