package libserve

import (
	"bytes"
	"io"
	"net/http"
)

func NewCacher() Cacher {
	return Cacher{
		items: make(map[string]*http.Response),
	}
}

type Cacher struct {
	items map[string]*http.Response
}

func (c *Cacher) Has(path string) bool {
	_, ok := c.items[path]
	return ok
}

func (c *Cacher) Save(path string, res *http.Response) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	c.items[path] = &http.Response{
		Status:        res.Status,
		StatusCode:    res.StatusCode,
		Header:        res.Header,
		Body:          io.NopCloser(bytes.NewReader(body)),
		ContentLength: res.ContentLength,
	}
}

func (c *Cacher) Get(path string) *http.Response {
	return c.items[path]
}
