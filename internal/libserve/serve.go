package libserve

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/enuesaa/cywagon/internal/ctlconf"
)

type ServeConf struct {
	entryUrl *url.URL
	conf ctlconf.Conf
}
type ServeMap map[string]ServeConf

func Serve(confs []ctlconf.Conf) error {
	servemap := make(ServeMap)

	for i, conf := range confs {
		entryUrl, err := url.Parse(conf.Entry.Host)
		if err != nil {
			return err
		}
		servemap[conf.Host] = ServeConf{entryUrl: entryUrl, conf: conf}
		if i == 0 {
			servemap["default"] = ServeConf{entryUrl: entryUrl, conf: conf}
		}
	}
	proxy := httputil.ReverseProxy{}
	proxy.Rewrite = func(req *httputil.ProxyRequest) {
		serveConf, ok := servemap[req.In.Host]
		if !ok {
			serveConf, ok = servemap["default"]
			if !ok {
				return
			}
		}
		req.SetURL(serveConf.entryUrl)
	}
	proxy.Transport = &Transport{ServeMap: servemap}
	// proxy.ModifyResponse = func(resp *http.Response) error {
	// 	return nil
	// }
	// proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {}

	return http.ListenAndServe(":3000", &proxy)
}
