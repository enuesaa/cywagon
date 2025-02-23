package libfetch

import "github.com/enuesaa/cywagon/internal/infra"

func New(container infra.Container) Fetcher {
	return Fetcher{
		Container: container,
	}
}

type Fetcher struct {
	infra.Container
}
