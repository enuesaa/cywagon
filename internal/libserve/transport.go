package libserve

import "net/http"

// see https://engineering.mercari.com/blog/entry/2018-12-05-105737/
type Transport struct {
	ServeOpts ServeOpts
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	site := t.ServeOpts.getByHost(req.Host)

	next := func(r *http.Request) *http.Response {
		res, _ := http.DefaultTransport.RoundTrip(r)
		return res
	}
	return site.Handler(req, next)
}
