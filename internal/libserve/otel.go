package libserve

import (
	"context"
	"net/http"
)

type OtelMiddleware struct {
	Next Middleware
}

func (m *OtelMiddleware) Handle(site Site, req *http.Request) (*http.Response, error) {
	if tracer == nil {
		return m.Next.Handle(site, req)
	}

	_, span := tracer.Start(context.Background(), "handle")
	defer span.End()

	return m.Next.Handle(site, req)
}
