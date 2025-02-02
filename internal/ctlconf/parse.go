package ctlconf

import "github.com/enuesaa/cywagon/internal/liblua"

func parse(code string) (Conf, error) {
	config := Conf{
		Host: "aa",
		Entry: ConfEntry{
			Workdir:        ".",
			Cmd:            "",
			WaitForHealthy: 60,
		},
	}
	runner := liblua.NewRunner(code)

	if err := runner.Blend(config); err != nil {
		return config, err
	}

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
	config.Handler = runner.GetFunction("handler")

	if err := runner.GetTable("entry", &entry); err != nil {
		return config, err
	}
	config.Entry = entry

	return config, nil
}
