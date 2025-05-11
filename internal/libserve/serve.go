package libserve

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func (s *Server) Use(handler Handler) {
	s.handlers = append(s.handlers, handler)
}

func (s *Server) UseTLS(certFile string, keyFile string) {
	s.certs = append(s.certs, Cert{certFile: certFile, keyFile: keyFile})
}

func (s *Server) Serve() error {
	addr := fmt.Sprintf(":%d", s.Port)

	return http.ListenAndServe(addr, s)
}

func (s *Server) ServeTLS() error {
	addr := fmt.Sprintf(":%d", s.Port)

	// see https://gist.github.com/denji/12b3a568f092ab951456
	tlsconfig := tls.Config{
		Certificates: make([]tls.Certificate, 0),
	}
	for _, c := range s.certs {
		cert, err := tls.LoadX509KeyPair(c.certFile, c.keyFile)
		if err != nil {
			return err
		}
		tlsconfig.Certificates = append(tlsconfig.Certificates, cert)
	}
	srv := &http.Server{
		Addr:      addr,
		Handler:   s,
		TLSConfig: &tlsconfig,
	}
	return srv.ListenAndServeTLS("", "")
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := NewContext(req)

	for _, handler := range s.handlers {
		res := handler(&ctx)
		if res != nil {
			s.OnResponse(&ctx, res.status, ctx.req.Method)

			if err := res.flush(w); err != nil {
				s.OnError(&ctx, err)
			}
			break
		}
	}
}
