package libserve

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/enuesaa/cywagon/internal/ctlconf"
)

func Serve(confs []ctlconf.Conf) error {
	parsedHost, err := url.Parse(confs[0].Entry.Host)
	if err != nil {
		return err
	}

	proxy := httputil.ReverseProxy{}
	proxy.Rewrite = func(req *httputil.ProxyRequest) {
		req.SetURL(parsedHost)
		// req.Out.Host = parsedHost.Host
	}
	proxy.Transport = &Transport{}
	// proxy.ModifyResponse = func(resp *http.Response) error {
	// 	return nil
	// }
	// proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {}

	return http.ListenAndServe(":3000", &proxy)
}
