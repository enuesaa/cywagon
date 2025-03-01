package libserve

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type CacheMiddleware struct {
	cache map[string]HttpCache
	Next Middleware
}

type HttpCache struct {
	res *http.Response
	body bytes.Buffer
}

func (m *CacheMiddleware) key(req *http.Request) string {
	return fmt.Sprintf("%s/%s", req.Host, req.URL.Path)
}

func (m *CacheMiddleware) get(req *http.Request) (*http.Response, bool) {
	key := m.key(req)
	item, ok := m.cache[key]
	if !ok {
		return nil, false
	}
	res := item.res
	res.Body = io.NopCloser(&item.body)

	return res, true
}

func (m *CacheMiddleware) save(req *http.Request, res *http.Response, resbody bytes.Buffer) {
	key := m.key(req)
	m.cache[key] = HttpCache{
		res: &http.Response{
			Status:        res.Status,
			StatusCode:    res.StatusCode,
			Header:        res.Header,
			ContentLength: res.ContentLength,
		},
		body: resbody,
	}
}

func (m *CacheMiddleware) Handle(site Site, req *http.Request) (*http.Response, error) {
	if !site.Cache {
		return m.Next.Handle(site, req)
	}

	if res, ok := m.get(req); ok {
		fmt.Printf("use cache\n")
		return res, nil
	}

	// next
	res, err := m.Next.Handle(site, req)
	if err != nil {
		return nil, err
	}

	var resbody bytes.Buffer

	if _, err := io.Copy(&resbody, res.Body); err != nil {
		return res, err
	}
	defer res.Body.Close()

	m.save(req, res, resbody)
	res.Body = io.NopCloser(&resbody)

	return res, nil
}
