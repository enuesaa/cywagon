package libserve

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/enuesaa/cywagon/internal/ctlconf"
)

func Serve(confs []ctlconf.Conf) error {
	servemap := map[string]*url.URL{}

	for i, conf := range confs {
		entryUrl, err := url.Parse(confs[0].Entry.Host)
		if err != nil {
			return err
		}
		servemap[conf.Host] = entryUrl
		if i == 0 {
			servemap["default"] = entryUrl
		}
	}
	proxy := httputil.ReverseProxy{}
	proxy.Rewrite = func(req *httputil.ProxyRequest) {
		entryUrl, ok := servemap[req.In.Host]
		if !ok {
			entryUrl, ok = servemap["default"]
			if !ok {
				return
			}
		}
		req.SetURL(entryUrl)
	}
	// proxy.ModifyResponse = func(resp *http.Response) error {
	// 	return nil
	// }
	// proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {}

	return http.ListenAndServe(":3000", &proxy)
}
