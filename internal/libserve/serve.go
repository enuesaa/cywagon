package libserve

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func (s *Server) Use(handler Handler) {
	s.handlers = append(s.handlers, handler)
}

func (s *Server) Listen(port int) {
	s.listeners = append(s.listeners, Listener{
		port: port,
		tls: false,
	})
}

func (s *Server) ListenTLS(port int, certfile string, keyfile string) {
	s.listeners = append(s.listeners, Listener{
		port: port,
		tls: false,
		certfile: certfile,
		keyfile: keyfile,
	})
}

func (s *Server) Serve() error {
	addr := fmt.Sprintf(":%d", s.Port)
	listener := Listener{
		server: s,
		port: s.Port,
		tls: false,
	}
	srv := &http.Server{
		Addr:    addr,
		Handler: &listener,
	}
	return srv.ListenAndServe()
}

func (s *Server) ServeTLS() error {
	addr := fmt.Sprintf(":%d", s.Port)

	// see https://gist.github.com/denji/12b3a568f092ab951456
	tlsconfig := tls.Config{
		Certificates: make([]tls.Certificate, 0),
	}
	for _, l := range s.listeners {
		cert, err := tls.LoadX509KeyPair(l.certfile, l.keyfile)
		if err != nil {
			return err
		}
		tlsconfig.Certificates = append(tlsconfig.Certificates, cert)
	}
	listener := Listener{
		server: s,
		port: s.Port,
		tls: false,
	}
	srv := &http.Server{
		Addr:      addr,
		Handler:   &listener,
		TLSConfig: &tlsconfig,
	}
	return srv.ListenAndServeTLS("", "")
}
