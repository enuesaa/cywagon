package libserve

import (
	"io"
	"mime"
	"net/http"
	"path/filepath"
)

type Context struct {
	Id           string
	Port         int
	Host         string
	Path         string
	Headers      map[string]string
	res          Response
	req          *http.Request
	statusPrefer int
}

func (c *Context) ResHeader(name string, value string) {
	c.res.headers[name] = value
}

func (c *Context) ResBody(path string, body io.Reader) error {
	contentType := c.CalcContentType(path)
	if contentType != "" {
		c.res.headers["Content-Type"] = contentType
	}

	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	c.res.body = b

	return nil
}

func (c *Context) CalcContentType(path string) string {
	ext := filepath.Ext(path)
	return mime.TypeByExtension(ext)
}

func (c *Context) ResStatusPrefer(status int) {
	c.statusPrefer = status
}

func (c *Context) Resolve(status int) *Response {
	if c.statusPrefer > 0 {
		c.res.status = c.statusPrefer
	} else {
		c.res.status = status
	}
	return &c.res
}

func (c *Context) GetOriginalPath() string {
	return c.req.URL.Path
}
