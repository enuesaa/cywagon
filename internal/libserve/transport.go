package libserve

import "net/http"

type Transport struct {
	Sites Sites
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	var res *http.Response

	site := t.Sites.getByHost(req.Host)

	next := func(rq HandlerRequest) HandlerResponse {
		req.URL.Path = rq.Path

		res, _ = http.DefaultTransport.RoundTrip(req)

		rs := HandlerResponse{
			Status: res.StatusCode,
		}
		return rs
	}
	rq := HandlerRequest{
		Path: req.URL.Path,
	}
	rs := HandlerResponse{
		Status: 0,
	}
	if err := site.Handler(&rs, next, rq); err != nil {
		return res, err
	}

	res.StatusCode = rs.Status

	return res, nil
}
