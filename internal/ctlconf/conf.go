package ctlconf

import (
	"net/http"

	"github.com/enuesaa/cywagon/internal/liblua"
)

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

func (c *Conf) RunHandler(realnext func() *http.Response) (int, error) {
	type Request struct{}
	type Response struct {
		Status int `lua:"status"`
	}
	var req Request
	var res Response

	next := func(req interface{}) interface{} {
		httpres := realnext()
		res.Status = httpres.StatusCode
		return res
	}
	luaval, err := c.Handler(next, req)
	if err != nil {
		return 0, err
	}
	if err := liblua.Unmarshal(luaval, &res); err != nil {
		return 0, err
	}
	return res.Status, nil
}
