package libserve

import "net/http"

type HandleMiddleware struct {}

func (m *HandleMiddleware) Handle(site Site, req *http.Request) (*http.Response, error) {
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
	if err := site.Handler(&rs, next, rq); err != nil {
		return res, err
	}
	res.StatusCode = rs.Status

	return res, nil
}
