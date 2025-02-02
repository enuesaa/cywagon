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
		HealthCheck: ConfHealthCheck{
			Protocol: "HTTP",
			Method:   "GET",
			Path:     "/",
		},
		Handler: liblua.Fn{},
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
