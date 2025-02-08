package libserve

import "net/http"

// see https://engineering.mercari.com/blog/entry/2018-12-05-105737/
type Transport struct {
	ServeMap ServeMap
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	var res *http.Response
	var err error

	serveConf := t.ServeMap.Get(req.Host)

	next := func() *http.Response {
		res, _ = http.DefaultTransport.RoundTrip(req)
		return res
	}
	code, _ := serveConf.conf.RunHandler(next)
	res.StatusCode = code

	return res, err
}
