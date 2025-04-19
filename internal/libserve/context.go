package libserve

import (
	"io"
	"net/http"
)

func NewContext(req *http.Request) Context {
	return Context{
		Host: req.Host,
		Path: req.URL.Path,
		req: req,
		resHeaders: make(map[string]string),
		resBody: nil,
		resPath: "",
	}
}

type Context struct {
	Host string
	Path string
	req *http.Request
	resHeaders map[string]string
	resBody    io.Reader
	resPath    string
}

func (c *Context) SetResponseHeader(name, value string) {
	c.resHeaders[name] = value
}

func (c *Context) SetResponseBody(path string, body io.Reader) {
	c.resPath = path
	c.resBody = body
}

func (c *Context) Resolve(status int) *Response {
	return &Response{
		headers: c.resHeaders,
		status: status,
		body: c.resBody,
		path: c.resPath,
	}
}
