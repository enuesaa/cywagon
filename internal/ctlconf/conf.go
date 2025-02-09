package ctlconf

import (
	"net/http"

	"github.com/enuesaa/cywagon/internal/liblua"
	"github.com/enuesaa/cywagon/internal/libserve"
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

// libserve に移したい
type ConfHandlerRequest struct{}
type ConfHandlerResponse struct {
	Status int `lua:"status"`
}

func (c *Conf) RunHandler(req *http.Request, next libserve.FnNext, res *http.Response) error {
	handlerReq := ConfHandlerRequest{}
	handlerRes := ConfHandlerResponse{
		Status: 0,
	}

	args := []interface{}{}
	args = append(args, func(ConfHandlerRequest) ConfHandlerResponse {
		res = next(req)
		handlerRes.Status = res.StatusCode
		return handlerRes
	})
	args = append(args, handlerReq)

	if err := c.Handler(args, &handlerRes); err != nil {
		return err
	}
	res.StatusCode = handlerRes.Status

	return nil
}
