package enginectl

import "github.com/enuesaa/cywagon/internal/libfetch"

func (e *Engine) CheckHealth() error {
	return libfetch.Fetch()
}
