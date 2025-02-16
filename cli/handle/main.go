package handle

import "github.com/enuesaa/cywagon/internal/infra"

func New() Handler {
	return Handler{
		Container: infra.Default,
	}
}

type Handler struct {
	infra.Container
}
