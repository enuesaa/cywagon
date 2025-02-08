package libserve

import (
	"fmt"
	"net/http"
)

// see https://engineering.mercari.com/blog/entry/2018-12-05-105737/
// proxy.Transport
type Transport struct {
	ServeMap ServeMap
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	var res *http.Response
	var err error

	serveConf, ok := t.ServeMap[req.Host]
	if !ok {
		serveConf, ok = t.ServeMap["default"]
		if !ok {
			return res, err
		}
	}

	next := func() *http.Response {
		fmt.Printf("next func\n")
		res, _ = http.DefaultTransport.RoundTrip(req)
		return res
	}
	code, _ := serveConf.conf.RunHandler(next)
	fmt.Printf("code: %d\n", code)

	res.StatusCode = code

	return res, err
}
