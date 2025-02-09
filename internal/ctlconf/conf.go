package ctlconf

import (
	"github.com/enuesaa/cywagon/internal/liblua"
	"github.com/enuesaa/cywagon/internal/libserve"
)

type Conf struct {
	Host        string          `lua:"host"`
	Entry       ConfEntry       `lua:"entry"`
	HealthCheck ConfHealthCheck `lua:"healthCheck"`
	Handler     liblua.FnI      `lua:"handler"`
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
	next := func(req ConfHandlerRequest) ConfHandlerResponse {
		serveRes := serveNext(nil)
		res := ConfHandlerResponse{
			Status: serveRes.StatusCode,
		}
		return res
	}
	req := ConfHandlerRequest{}
	args := []interface{}{next, req}
	
	// TODO
	c.Handler(args)

	return nil
}
