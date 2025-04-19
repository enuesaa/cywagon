package libserve

import "net/http"

type Site struct {
	Host string // Example: `example.com`
	Handle func(w http.ResponseWriter, req *http.Request)
}
