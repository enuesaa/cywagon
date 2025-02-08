package ctlconf

import (
	"context"

	"github.com/enuesaa/cywagon/internal/liblua"
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

func parse(code string) (Conf, error) {
	config := Conf{
		Host: "aa",
		Entry: ConfEntry{
			Workdir:        ".",
			Cmd:            "",
			WaitForHealthy: 60,
		},
		HealthCheck: ConfHealthCheck{
			Protocol: "HTTP",
			Method:   "GET",
			Path:     "/",
		},
		Handler: nil,
	}
	runner := liblua.NewRunner(code)

	if err := runner.Inject(config); err != nil {
		return config, err
	}
	if err := runner.Run(); err != nil {
		return config, err
	}
	if err := runner.Eject(&config); err != nil {
		return config, err
	}
	return config, nil
}
