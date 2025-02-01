package libserve

import "net/http"

// see https://engineering.mercari.com/blog/entry/2018-12-05-105737/

type Transport struct {
	Host string
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Host = t.Host

	return http.DefaultTransport.RoundTrip(req)
}
