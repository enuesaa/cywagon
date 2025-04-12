package libserve

import (
	"fmt"
	"net/http"
)

// see https://engineering.mercari.com/blog/entry/2018-12-05-105737/
// also see ModifyResponse, ErrorHandler if need

func (s *Server) Serve() error {
	if err := s.Sites.Validate(); err != nil {
		return err
	}
	handler := FsHandler{
		Sites: s.Sites,
	}
	addr := fmt.Sprintf(":%d", s.Port)

	return http.ListenAndServe(addr, &handler)
}

type Middleware interface {
	Handle(site Site, req *http.Request) (*http.Response, error)
}

type FsHandler struct {
	Sites Sites
}

func (h *FsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	host := req.Host

	site := h.Sites.getByHost(host)
	http.ServeFileFS(w, req, site.Dist, ".")
}
