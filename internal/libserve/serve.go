package libserve

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func (s *Server) Use(handler Handler) {
	s.handlers = append(s.handlers, handler)
}

func (s *Server) Listen(port int) {
	if _, ok := s.listenmap[port]; !ok {
		s.listenmap[port] = make([]ListenConfig, 0)
	}
	s.listenmap[port] = append(s.listenmap[port], ListenConfig{
		tls: false,
	})
}

func (s *Server) ListenTLS(port int, certfile string, keyfile string) {
	if _, ok := s.listenmap[port]; !ok {
		s.listenmap[port] = make([]ListenConfig, 0)
	}
	s.listenmap[port] = append(s.listenmap[port], ListenConfig{
		tls: true,
		certfile: certfile,
		keyfile: keyfile,
	})
}

func (s *Server) Serve() error {
	g, _ := errgroup.WithContext(context.Background())

	for port, config := range s.listenmap {
		isTls, err := s.judgeIsTLSListener(config)
		if err != nil {
			return err
		}
		if isTls {
			g.Go(s.serveTLS(port, config))
		} else {
			g.Go(s.serve(port))
		}
	}
	return g.Wait()
}

func (s *Server) judgeIsTLSListener(config []ListenConfig) (bool, error) {
	tls := config[0].tls
	for _, conf := range config {
		if conf.tls != tls {
			return false, fmt.Errorf("invalid port mapping")
		}
	}
	return tls, nil
}

func (s *Server) serve(port int) func() error {
	return func() error {
		addr := fmt.Sprintf(":%d", port)
		listener := Listener{
			Server: s,
			port: port,
		}
		srv := &http.Server{
			Addr:    addr,
			Handler: &listener,
		}
		return srv.ListenAndServe()
	}
}

func (s *Server) serveTLS(port int, lconfig []ListenConfig) func() error {
	return func() error {
		addr := fmt.Sprintf(":%d", port)

		// see https://gist.github.com/denji/12b3a568f092ab951456
		tlsconfig := tls.Config{
			Certificates: make([]tls.Certificate, 0),
		}
		for _, l := range lconfig {
			cert, err := tls.LoadX509KeyPair(l.certfile, l.keyfile)
			if err != nil {
				return err
			}
			tlsconfig.Certificates = append(tlsconfig.Certificates, cert)
		}
		listener := Listener{
			Server: s,
			port: port,
		}
		srv := &http.Server{
			Addr:      addr,
			Handler:   &listener,
			TLSConfig: &tlsconfig,
		}
		return srv.ListenAndServeTLS("", "")
	}
}
