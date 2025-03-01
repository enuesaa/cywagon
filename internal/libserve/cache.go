package libserve

import (
	"bytes"
	"io"
	"net/http"
)

func NewCacher() Cacher {
	return Cacher{
		items: make(map[string]*http.Response),
		body: make(map[string]bytes.Buffer),
	}
}

type Cacher struct {
	items map[string]*http.Response
	body map[string]bytes.Buffer
}

func (c *Cacher) Has(path string) bool {
	_, ok := c.items[path]
	return ok
}

func (c *Cacher) Save(path string, res *http.Response, resbody bytes.Buffer) {
	c.items[path] = &http.Response{
		Status:        res.Status,
		StatusCode:    res.StatusCode,
		Header:        res.Header,
		// Body:          io.NopCloser(bytes.NewReader(body)),
		ContentLength: res.ContentLength,
	}
	c.body[path] = resbody
}

func (c *Cacher) Get(path string) *http.Response {
	res := c.items[path]
	body := c.body[path]

	res.Body = io.NopCloser(&body)

	return res
}
