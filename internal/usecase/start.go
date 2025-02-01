package usecase

import (
	"context"

	"github.com/enuesaa/cywagon/internal/libserve"
)

func Start(ctx context.Context, confDir string) error {
	return libserve.Serve()
}
