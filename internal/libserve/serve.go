package libserve

import (
	"net/http"
	"net/http/httputil"

	"github.com/enuesaa/cywagon/internal/ctlconf"
)

func Serve(confs []ctlconf.Conf) error {
	serveMap := newServeMap(confs)

	proxy := httputil.ReverseProxy{}
	proxy.Rewrite = func(req *httputil.ProxyRequest) {
		serveConf := serveMap.Get(req.In.Host)
		req.SetURL(serveConf.entryUrl)
	}
	proxy.Transport = &Transport{ServeMap: serveMap}
	// proxy.ModifyResponse = func(resp *http.Response) error {
	// 	return nil
	// }
	// proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {}

	return http.ListenAndServe(":3000", &proxy)
}
