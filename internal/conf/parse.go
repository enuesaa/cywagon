package conf

import (
	"context"

	"github.com/enuesaa/cywagon/internal/liblua"
	"github.com/enuesaa/cywagon/internal/repository"
)

func Parse(ctx context.Context, path string) (Config, error) {
	repos := repository.Use(ctx)

	var config Config

	scriptbytes, err := repos.Fs.Read(path)
	if err != nil {
		return config, err
	}

	runner := liblua.NewRunner(string(scriptbytes))
	if err := runner.Run(); err != nil {
		return config, err
	}
	config.Hostname = runner.GetString("hostname")
	config.Port = runner.GetInt("port")
	config.handler = runner.GetFunction("handler")

	return config, nil
}
