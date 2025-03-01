package model

import (
	"fmt"
	"net/url"

	"github.com/enuesaa/cywagon/internal/liblua"
)

type Conf struct {
	Host        string          `lua:"host"`
	Origin      ConfOrigin      `lua:"origin"`
	HealthCheck ConfHealthCheck `lua:"healthCheck"`
	Handler     liblua.Fn       `lua:"handler"`
	Cache       bool            `lua:"cache"`
}

type ConfOrigin struct {
	Url            string `lua:"url"`
	Workdir        string `lua:"workdir"`
	Cmd            string `lua:"cmd"`
	WaitForHealthy int    `lua:"waitForHealthy"`
}

type ConfHealthCheck struct {
	Protocol string `lua:"protocol"`
	Method   string `lua:"method"`
	Path     string `lua:"path"`
	Matcher  string `lua:"matcher"`
}

func (c *Conf) HealthCheckUrl() string {
	u, err := url.Parse(c.Origin.Url)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s%s", u.Host, c.HealthCheck.Path)
}
