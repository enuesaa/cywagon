package enginectl

import "github.com/enuesaa/cywagon/internal/libfetch"

func (e *Engine) CheckHealth() error {
	fetcher := libfetch.New()

	return fetcher.CheckTcpConn()
}
