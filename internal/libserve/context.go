package libserve

import (
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

func NewContext(req *http.Request) Context {
	headers := make(map[string]string)
	for key, value := range req.Header {
		headers[key] = value[0]
	}

	return Context{
		Host: req.Host,
		Path: req.URL.Path,
		Headers: headers,
		req: req,
		resHeaders: make(map[string]string),
	}
}

type Context struct {
	Host string
	Path string
	Headers map[string]string
	req *http.Request
	resHeaders map[string]string
	resBody    []byte
}

func (c *Context) GetLookupPath() string {
	path := c.Path
	if strings.HasSuffix(path, "/") {
		path = filepath.Join(path, "index.html")
	}
	return strings.TrimPrefix(path, "/")
}

func (c *Context) SetResponseHeader(name, value string) {
	c.resHeaders[name] = value
}

func (c *Context) SetResponseBody(path string, body io.Reader) error {
	ext := filepath.Ext(path)
	c.resHeaders["Content-Type"] = mime.TypeByExtension(ext)

	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	c.resBody = b

	return nil
}

func (c *Context) Resolve(status int) *Response {
	return &Response{
		headers: c.resHeaders,
		status: status,
		body: c.resBody,
	}
}
