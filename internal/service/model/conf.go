package model

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
type ConfHandlerRequest struct{
	Path string `lua:"path"`
}
type ConfHandlerResponse struct {
	Status int `lua:"status"`
}

func (c *Conf) RunHandler(req *http.Request, next libserve.FnNext) (*http.Response, error) {
	var res *http.Response

	handlerReq := ConfHandlerRequest{
		Path: req.URL.Path,
	}
	handlerRes := ConfHandlerResponse{
		Status: 0,
	}
	handlerNext := func(r ConfHandlerRequest) ConfHandlerResponse {
		req.URL.Path = r.Path
		res = next(req)
		handlerRes.Status = res.StatusCode
		return handlerRes
	}

	if err := c.Handler(&handlerRes, handlerNext, handlerReq); err != nil {
		return res, err
	}
	res.StatusCode = handlerRes.Status

	return res, nil
}
