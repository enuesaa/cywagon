package ctlconf

import (
	"net/http"

	"github.com/enuesaa/cywagon/internal/liblua"
	lua "github.com/yuin/gopher-lua"
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

	next := func(luareq *lua.LTable) *lua.LTable {
		httpres := realnext()
		res := Response{
			Status: httpres.StatusCode,
		}
		luares, _ := liblua.Marshal(res)
		return luares
	}
	res, err := c.Handler(next, req)
	if err != nil {
		return 0, err
	}
	status := res.GetInt("status")

	return status, nil
}
