package ctlconf

import (
	"fmt"
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

type ConfHandlerRequest struct{}
type ConfHandlerResponse struct{
	Status int `lua:"status"`
}

//
func (c *Conf) RunHandler(req *http.Request, next libserve.FnNext) *http.Response {
	var httpres *http.Response

	args := []interface{}{}
	args = append(args, ConfHandlerRequest{})
	args = append(args, func(ConfHandlerRequest) ConfHandlerResponse {
		httpres = next(req)
		return ConfHandlerResponse{ Status: 200 }
	})
	res := ConfHandlerResponse{
		Status: 200,
	}

	if err := c.Handler(args, &res); err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Printf("%+v\n", res)

	httpres.StatusCode = res.Status

	return httpres
}
