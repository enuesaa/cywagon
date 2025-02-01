package ctlconf

import (
	"fmt"

	"github.com/enuesaa/cywagon/internal/liblua"
)

type Config struct {
	Hostname    string
	Entry       ConfigEntry
	HealthCheck ConfigHealthCheck
	handler     liblua.Fn
}
type ConfigEntry struct {
	Workdir        string `lua:"workdir"`
	Cmd            string `lua:"cmd"`
	WaitForHealthy int    `lua:"waitForHealthy"`
}
type ConfigHealthCheck struct {
	Protocol string `lua:"protocol"`
	Method   string `lua:"method"`
	Path     string `lua:"path"`
}

func (c *Config) RunHandler() error {
	type Response struct {
		Status int `lua:"status"`
	}
	response := Response{
		Status: 404,
	}
	next := func() {
		fmt.Println("this is next function")
	}

	result, err := c.handler.Run(next, nil, response)
	if err != nil {
		return err
	}
	fmt.Printf("res: %d\n", result.GetInt("status"))

	return nil
}
