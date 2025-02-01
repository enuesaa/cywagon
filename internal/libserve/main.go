package libserve

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Serve(uri string) error {
	parsedHost, err := url.Parse(uri)
	if err != nil {
		return err
	}

	proxy := httputil.NewSingleHostReverseProxy(parsedHost)

	proxy.Transport = &Transport{
		Host: parsedHost.Host,
	}
	// proxy.ModifyResponse = func(resp *http.Response) error {
	// 	return nil
	// }
	// proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {}

	return http.ListenAndServe(":3000", proxy)
}
