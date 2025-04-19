package libserve

import "net/http"

type Handler struct {
	sites map[string]Site
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	host := req.Host

	site := h.getByHost(host)
	site.Handle(w, req)
}

func (h *Handler) getByHost(host string) Site {
	site, ok := h.sites[host]
	if ok {
		return site
	}
	return h.sites["default"]
}
