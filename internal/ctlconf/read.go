package ctlconf

import (
	"context"

	"github.com/enuesaa/cywagon/internal/repository"
)

func Read(ctx context.Context, path string) (Conf, error) {
	repos := repository.Use(ctx)

	codeb, err := repos.Fs.Read(path)
	if err != nil {
		return Conf{}, err
	}
	code := string(codeb)

	return parse(code)
}
