package libserve

import (
	"net/http"

	"github.com/enuesaa/cywagon/internal/infra"
)

type Transport struct {
	infra.Container
	Sites Sites
	Cacher Cacher
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	var res *http.Response

	site := t.Sites.getByHost(req.Host)

	if t.Cacher.Has(req.URL.Path) {
		res = t.Cacher.Get(req.URL.Path)
		t.Log.Info("[http] %d %s %s %s (cache)", res.StatusCode, req.Method, site.Host, req.URL.Path)
		return res, nil
	}

	rq := HandlerRequest{
		Path: req.URL.Path,
	}
	rs := HandlerResponse{
		Status: 0,
	}
	next := func(rq HandlerRequest) HandlerResponse {
		req.URL.Path = rq.Path
		// https://seomis.cc/blog/content-encoding-golang-transport
		req.Header.Set("Accept-Encoding", "deflate, br, zstd")

		res, _ = http.DefaultTransport.RoundTrip(req)
		rs.Status = res.StatusCode

		return rs
	}

	if err := site.Handler(&rs, next, rq); err != nil {
		return res, err
	}
	res.StatusCode = rs.Status

	t.Cacher.Save(req.URL.Path, res)

	t.Log.Info("[http] %d %s %s %s", res.StatusCode, req.Method, site.Host, req.URL.Path)

	return res, nil
}
