package libserve

import "net/http"

type Middleware interface {
	Handle(site Site, req *http.Request) (*http.Response, error)
}
