package libserve

import (
	"net/http"
	"net/http/httputil"
)

// see https://engineering.mercari.com/blog/entry/2018-12-05-105737/
// also see ModifyResponse, ErrorHandler if need

func (s *Server) Serve() error {
	proxy := httputil.ReverseProxy{}
	proxy.Rewrite = func(req *httputil.ProxyRequest) {
		site := s.Sites.getByHost(req.In.Host)
		req.SetURL(site.parsedOriginUrl)
	}
	proxy.Transport = &Transport{
		Sites: s.Sites,
	}

	return http.ListenAndServe(":3000", &proxy)
}
