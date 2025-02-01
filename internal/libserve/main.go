package libserve

import (
	"fmt"

	"net/http"
	"net/http/httputil"
	"net/url"
)

// see https://engineering.mercari.com/blog/entry/2018-12-05-105737/
type CustomTransport struct{}

func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Host = "example.com"
	fmt.Printf("%+v\n", req)

	return http.DefaultTransport.RoundTrip(req)
}

func Serve() error {
	host, err := url.Parse("https://example.com")
	if err != nil {
		return err
	}

	proxy := httputil.NewSingleHostReverseProxy(host)

	proxy.Transport = &CustomTransport{}
	proxy.ModifyResponse = func(resp *http.Response) error {
		return nil
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {}

	return http.ListenAndServe(":3000", proxy)
}
