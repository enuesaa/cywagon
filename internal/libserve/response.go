package libserve

import "net/http"

type Response struct {
	headers map[string]string
	status int
	body []byte
}

func (r *Response) flush(w http.ResponseWriter) error {
	for name, value := range r.headers {
		w.Header().Set(name, value)
	}
	w.WriteHeader(r.status)

	if r.body != nil {
		if _, err := w.Write(r.body); err != nil {
			return err
		}
	}
	return nil
}
