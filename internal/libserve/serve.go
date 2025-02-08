package libserve

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/enuesaa/cywagon/internal/ctlconf"
)

type ServeConf struct {
	EntryUrl *url.URL
}

func Serve(confs []ctlconf.Conf) error {
	servemap := map[string]ServeConf{}

	for _, conf := range confs {
		entryUrl, err := url.Parse(confs[0].Entry.Host)
		if err != nil {
			return err
		}
		servemap[conf.Host] = ServeConf{
			EntryUrl: entryUrl,
		}
	}
	proxy := httputil.ReverseProxy{}
	proxy.Rewrite = func(req *httputil.ProxyRequest) {
		serveConf, ok := servemap[req.In.Host]
		if !ok {
			return
		}
		req.SetURL(serveConf.EntryUrl)
	}
	// proxy.ModifyResponse = func(resp *http.Response) error {
	// 	return nil
	// }
	// proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {}

	return http.ListenAndServe(":3000", &proxy)
}
