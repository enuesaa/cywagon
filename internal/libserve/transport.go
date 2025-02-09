package libserve

import "net/http"

// see https://engineering.mercari.com/blog/entry/2018-12-05-105737/
type Transport struct {
	ServeOpts ServeOpts
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	site := t.ServeOpts.getByHost(req.Host)
	next := func(req *http.Request) *http.Response {
		res, _ := http.DefaultTransport.RoundTrip(req)
		return res
	}
	res := site.Handler(req, next)

	return res, nil
}
