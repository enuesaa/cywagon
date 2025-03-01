package libserve

import (
	"bytes"
	"io"
	"net/http"
)

type HttpMiddleware interface {
	Handle(site Site, req *http.Request) (*http.Response, error)
}


type CacheMiddleware struct {
	Cacher Cacher
	Next HttpMiddleware
}

func (m *CacheMiddleware) Handle(site Site, req *http.Request) (*http.Response, error) {
	if m.Cacher.Has(req.URL.Path) {
		res := m.Cacher.Get(req.URL.Path)

		return res, nil
	}

	res, err := m.Next.Handle(site, req)
	if err != nil {
		return nil, err
	}

	var resbody bytes.Buffer

	if _, err := io.Copy(&resbody, res.Body); err != nil {
		return res, err
	}
	defer res.Body.Close()

	m.Cacher.Save(req.URL.Path, res, resbody)
	res.Body = io.NopCloser(&resbody)

	return res, nil
}


type HandleMiddleware struct {
	Cacher Cacher
}

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
