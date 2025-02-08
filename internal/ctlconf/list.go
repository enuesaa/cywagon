package ctlconf

import (
	"context"

	"github.com/enuesaa/cywagon/internal/repository"
)

func List(ctx context.Context, dir string) []string {
	repos := repository.Use(ctx)

	list, err := repos.Fs.ListFiles(dir)
	if err != nil {
		list = []string{}
	}

	return list
}