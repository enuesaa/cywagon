package ctlconf

import (
	"fmt"

	"github.com/enuesaa/cywagon/internal/liblua"
	"github.com/enuesaa/cywagon/internal/libserve"
)

type Conf struct {
	Host        string          `lua:"host"`
	Entry       ConfEntry       `lua:"entry"`
	HealthCheck ConfHealthCheck `lua:"healthCheck"`
	Handler     liblua.Fn `lua:"handler"`
	Invoke      liblua.Fn `lua:"invoke"`
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

type ConfHandlerRequest struct{}
type ConfHandlerResponse struct{
	Status int `lua:"status"`
}

func (c *Conf) RunHandler(serveNext libserve.FnNext) error {
	req := ConfHandlerRequest{}

	res := c.Handler(req)
	fmt.Println(res)
	return nil
}
