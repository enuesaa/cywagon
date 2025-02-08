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
	// TODO: fix
	var res *http.Response
	var err error

	serveConf, ok := t.ServeMap[req.Host]
	if !ok {
		serveConf, ok = t.ServeMap["default"]
		if !ok {
			return res, err
		}
	}

	next := func() {
		fmt.Printf("next func\n")
		res, err = http.DefaultTransport.RoundTrip(req)
	}
	code, _ := serveConf.conf.RunHandler(next)
	fmt.Printf("code: %d\n", code)

	return res, err
}
