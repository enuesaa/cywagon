package libserve

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
)

// see https://engineering.mercari.com/blog/entry/2018-12-05-105737/
// also see ModifyResponse, ErrorHandler if need

func (s *Server) Serve() error {
	if err := s.Sites.Validate(); err != nil {
		return err
	}
	if err := setupTracer(context.Background()); err != nil {
		return err
	}

	proxy := httputil.ReverseProxy{}
	proxy.Rewrite = func(req *httputil.ProxyRequest) {
		site := s.Sites.getByHost(req.In.Host)
		req.SetURL(site.parsedOriginUrl)
	}
	proxy.Transport = &Transport{
		Container: s.Container,
		Sites: s.Sites,
		Middleware: &OtelMiddleware{
			Next: &CacheMiddleware{
				cache: make(map[string]HttpCache),
				Next: &HandleMiddleware{},
			},
		},
	}
	addr := fmt.Sprintf(":%d", s.Port)

	return http.ListenAndServe(addr, &proxy)
}

type Middleware interface {
	Handle(site Site, req *http.Request) (*http.Response, error)
}
