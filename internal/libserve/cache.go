package libserve

import (
	"bytes"
	"io"
	"net/http"
)

func NewCacher() Cacher {
	return Cacher{
		items: make(map[string]HttpCache),
	}
}

type HttpCache struct {
	res *http.Response
	body bytes.Buffer
}

type Cacher struct {
	items map[string]HttpCache
}

func (c *Cacher) Has(path string) bool {
	_, ok := c.items[path]
	return ok
}

func (c *Cacher) Save(path string, res *http.Response, resbody bytes.Buffer) {
	c.items[path] = HttpCache{
		res: &http.Response{
			Status:        res.Status,
			StatusCode:    res.StatusCode,
			Header:        res.Header,
			ContentLength: res.ContentLength,
		},
		body: resbody,
	}
}

func (c *Cacher) Get(path string) *http.Response {
	item := c.items[path]
	res := item.res
	res.Body = io.NopCloser(&item.body)

	return res
}
