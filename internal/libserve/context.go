package libserve

import (
	"io"
	"mime"
	"net/http"
	"path/filepath"
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
		res: Response{
			headers: make(map[string]string),
			status:  0,
			body:    nil,
		},
		req: req,
	}
}

type Context struct {
	Host string
	Path string
	Headers map[string]string
	res Response
	req *http.Request
}

func (c *Context) ResHeader(name string, value string) {
	c.res.headers[name] = value
}

func (c *Context) ResBody(path string, body io.Reader) error {
	ext := filepath.Ext(path)
	contentType := mime.TypeByExtension(ext)
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

func (c *Context) Resolve(status int) *Response {
	c.res.status = status

	return &c.res
}
