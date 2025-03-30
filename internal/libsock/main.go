package libsock

import (
	"os"

	"github.com/enuesaa/cywagon/internal/infra"
)

func New() Sock {
	return Sock{
		Container: infra.Default,
		Path: "/tmp/cywagon.sock",
	}
}

type Sock struct {
	infra.Container

	Path string
}

func (e *Sock) Exists() bool {
	if _, err := os.Stat(e.Path); err == nil {
		return true
	}
	return false
}
