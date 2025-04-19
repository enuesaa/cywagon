package libserve

import "net/http"

type Site struct {
	Host string // Example: `example.com`

	// deprecated 
	Handle func(w http.ResponseWriter, req *http.Request)
}
