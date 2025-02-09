package libserve

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ServeOpts struct {
	Port    int
	Sites   []ServeOptsSite
	siteMap map[string]ServeOptsSite
}
type ServeOptsSite struct {
	Host            string // Example: `example.com`
	OriginUrl       string // Example: `https://example.com`
	Handler         FnHandler
	parsedOriginUrl *url.URL
}
type FnHandler func(*http.Request, FnNext) *http.Response
type FnNext func(*http.Request) *http.Response

func (o *ServeOpts) createSiteMap() error {
	o.siteMap = make(map[string]ServeOptsSite, len(o.Sites))

	for i, site := range o.Sites {
		parsed, err := url.Parse(site.OriginUrl)
		if err != nil {
			return err
		}
		site.parsedOriginUrl = parsed

		o.siteMap[site.Host] = site
		if i == 0 {
			o.siteMap["default"] = site
		}
	}
	return nil
}

func (m *ServeOpts) getByHost(host string) ServeOptsSite {
	site, ok := m.siteMap[host]
	if ok {
		return site
	}
	site, ok = m.siteMap["default"]
	if ok {
		return site
	}
	return ServeOptsSite{}
}

func Serve(opts ServeOpts) error {
	if err := opts.createSiteMap(); err != nil {
		return err
	}

	proxy := httputil.ReverseProxy{}
	proxy.Rewrite = func(req *httputil.ProxyRequest) {
		site := opts.getByHost(req.In.Host)
		req.SetURL(site.parsedOriginUrl)
	}
	proxy.Transport = &Transport{ServeOpts: opts}
	// proxy.ModifyResponse = func(resp *http.Response) error {
	// 	return nil
	// }
	// proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {}

	return http.ListenAndServe(":3000", &proxy)
}
