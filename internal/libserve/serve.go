package libserve

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func (o *Server) createSiteMap() (SiteMap, error) {
	sitemap := SiteMap{
		items: map[string]Site{},
	}

	for i, site := range o.Sites {
		parsed, err := url.Parse(site.OriginUrl)
		if err != nil {
			return sitemap, err
		}
		site.parsedOriginUrl = parsed

		sitemap.items[site.Host] = site
		if i == 0 {
			sitemap.items["default"] = site
		}
	}
	return sitemap, nil
}

func (s *Server) Serve() error {
	sitemap, err := s.createSiteMap()
	if err != nil {
		return err
	}

	proxy := httputil.ReverseProxy{}
	proxy.Rewrite = func(req *httputil.ProxyRequest) {
		site := sitemap.getByHost(req.In.Host)
		req.SetURL(site.parsedOriginUrl)
	}
	proxy.Transport = &Transport{SiteMap: sitemap}
	// also see ModifyResponse, ErrorHandler if need

	return http.ListenAndServe(":3000", &proxy)
}
