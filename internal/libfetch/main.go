package libfetch

import "github.com/enuesaa/cywagon/internal/infra"

func New() Fetcher {
	return Fetcher{
		Container: infra.Default,
	}
}

type Fetcher struct {
	infra.Container
}
