package libserve

import (
	"io"
	"net/http"
)

type Response struct {
	headers map[string]string
	status int
	body io.Reader
	path string
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
