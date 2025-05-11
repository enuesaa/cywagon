package libserve

import (
	"net/http"

	"github.com/oklog/ulid/v2"
)

type Listener struct {
	server *Server
	port int
	tls bool
	certfile string
	keyfile string
}

func (l *Listener) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	headers := make(map[string]string)
	for key, value := range req.Header {
		headers[key] = value[0]
	}

	ctx := Context{
		Id:      ulid.Make().String(),
		Host:    req.Host,
		Path:    req.URL.Path,
		Headers: headers,
		res: Response{
			headers: make(map[string]string),
			status:  0,
			body:    nil,
		},
		req:          req,
		statusPrefer: 0,
	}

	for _, handler := range l.server.handlers {
		res := handler(&ctx)
		if res != nil {
			l.server.OnResponse(&ctx, res.status, ctx.req.Method)

			if err := res.flush(w); err != nil {
				l.server.OnError(&ctx, err)
			}
			break
		}
	}
}
