package libserve

import "net/http"

// see https://engineering.mercari.com/blog/entry/2018-12-05-105737/
type Transport struct {
	SiteMap SiteMap
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	var res *http.Response

	site := t.SiteMap.getByHost(req.Host)

	next := func(rq FnHandlerRequest) FnHandlerResponse {
		req.URL.Path = rq.Path

		res, _ = http.DefaultTransport.RoundTrip(req)

		rs := FnHandlerResponse{
			Status: res.StatusCode,
		}
		return rs
	}
	rq := FnHandlerRequest{
		Path: req.URL.Path,
	}
	rs := FnHandlerResponse{
		Status: 0,
	}
	if err := site.Handler(&rs, next, rq); err != nil {
		return res, err
	}

	res.StatusCode = rs.Status

	return res, nil
}
