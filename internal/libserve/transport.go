package libserve

import (
	"net/http"

	"github.com/enuesaa/cywagon/internal/infra"
)

type Transport struct {
	infra.Container
	Sites Sites
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	var res *http.Response

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
	site := t.Sites.getByHost(req.Host)

	if err := site.Handler(&rs, next, rq); err != nil {
		return res, err
	}
	res.StatusCode = rs.Status

	t.Log.Info("[http] %d %s %s %s", res.StatusCode, req.Method, site.Host, req.URL.Path)

	return res, nil
}
