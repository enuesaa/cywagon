package libserve

import (
	"io"
	"net/http"
)

func NewResponse() Response {
	return Response{
		status: 0,
		body: nil,
		path: "",
		headers: make(map[string]string),
	}
}

type Response struct {
	headers map[string]string
	status int
	body io.Reader
	path string
}

func (r *Response) SetHeader(name string, value string) {
	r.headers[name] = value
}

func (r *Response) Write(status int, path string, body io.Reader) {
	r.status = status
	r.path = path
	r.body = body
}

func (r *Response) Close() bool {
	return r.status != 0
}

func (r *Response) Flush(w http.ResponseWriter) error {
	for name, value := range r.headers {
		w.Header().Set(name, value)
	}

	if r.body != nil {
		b, err := io.ReadAll(r.body)
		if err != nil {
			return err
		}
		if _, err := w.Write(b); err != nil {
			return err
		}
	}
	w.WriteHeader(r.status)

	return nil
}
