package libsock

import "github.com/enuesaa/cywagon/internal/infra"

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
