package libserve

import (
	"net/url"

	"github.com/enuesaa/cywagon/internal/ctlconf"
)

type ServeConf struct {
	entryUrl *url.URL
	conf     ctlconf.Conf
}
type ServeMap map[string]ServeConf

func (m *ServeMap) Get(host string) ServeConf {
	serveConf, ok := (*m)[host]
	if ok {
		return serveConf
	}
	serveConf, ok = (*m)["default"]
	if ok {
		return serveConf
	}
	return ServeConf{}
}

func newServeMap(confs []ctlconf.Conf) ServeMap {
	serveMap := make(ServeMap)

	for i, conf := range confs {
		entryUrl, err := url.Parse(conf.Entry.Host)
		if err != nil {
			continue
		}
		serveConf := ServeConf{
			entryUrl: entryUrl,
			conf:     conf,
		}
		serveMap[conf.Host] = serveConf

		if i == 0 {
			serveMap["default"] = serveConf
		}
	}
	return serveMap
}
