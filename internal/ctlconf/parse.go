package ctlconf

import (
	"context"

	"github.com/enuesaa/cywagon/internal/liblua"
)

func parse(ctx context.Context, code string) (Conf, error) {
	var config Conf
	runner := liblua.NewRunner(code)

	entry := ConfEntry{
		Workdir:        ".",
		Cmd:            "",
		WaitForHealthy: 60,
	}
	if err := runner.SetGlobal("entry", entry); err != nil {
		return config, err
	}
	healthCheck := ConfHealthCheck{
		Protocol: "HTTP",
		Method:   "GET",
		Path:     "/",
	}
	if err := runner.SetGlobal("healthCheck", healthCheck); err != nil {
		return config, err
	}

	if err := runner.Run(); err != nil {
		return config, err
	}
	config.Host = runner.GetString("host")
	config.handler = runner.GetFunction("handler")

	if err := runner.GetTable("entry", &entry); err != nil {
		return config, err
	}
	config.Entry = entry

	return config, nil
}
