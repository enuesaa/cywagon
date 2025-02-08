package ctlconf

import (
	"context"

	"github.com/enuesaa/cywagon/internal/liblua"
	"github.com/enuesaa/cywagon/internal/repository"
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

func (c *Conf) RunHandler(ctx context.Context) error {
	repos := repository.Use(ctx)

	type Response struct {
		Status int `lua:"status"`
	}
	response := Response{
		Status: 404,
	}
	next := func() {
		repos.Log.Info("this is next function")
	}

	result, err := c.Handler(next, nil, response)
	if err != nil {
		return err
	}
	repos.Log.Info("res: %d", result.GetInt("status"))

	return nil
}
