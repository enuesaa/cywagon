package libserve

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomTransport struct {}

func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Host = "example.com"
	fmt.Printf("%+v\n", req)

	return http.DefaultTransport.RoundTrip(req)
}

func Serve() error {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	exampleComUrl, err := url.Parse("https://example.com")
	if err != nil {
		return err
	}
	e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Skipper: func(e echo.Context) bool {
			return false
		},
		Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
			{
				URL: exampleComUrl,
			},
		}),
		Transport: &CustomTransport{},
	}))

	return e.Start(":3000")
}
