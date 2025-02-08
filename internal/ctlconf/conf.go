package ctlconf

import "github.com/enuesaa/cywagon/internal/liblua"

type Conf struct {
	Host        string          `lua:"host"`
	Entry       ConfEntry       `lua:"entry"`
	HealthCheck ConfHealthCheck `lua:"healthCheck"`
	Handler     liblua.Fn       `lua:"handler"`
}

type ConfEntry struct {
	Workdir        string `lua:"workdir"`
	Cmd            string `lua:"cmd"`
	WaitForHealthy int    `lua:"waitForHealthy"`
	Host           string `lua:"host"`
}

type ConfHealthCheck struct {
	Protocol string `lua:"protocol"`
	Method   string `lua:"method"`
	Path     string `lua:"path"`
}

func (c *Conf) RunHandler(next func()) (int, error) {
	type Response struct {
		Status int `lua:"status"`
	}
	response := Response{
		Status: 404,
	}
	result, err := c.Handler(next, nil, response)
	if err != nil {
		return 0, err
	}
	status := result.GetInt("status")

	return status, nil
}
