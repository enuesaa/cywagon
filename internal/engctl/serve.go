package engctl

import (
	"context"
	"net/http"

	"github.com/enuesaa/cywagon/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var data []byte

func Serve(ctx context.Context) error {
	repos := repository.Use(ctx)
	readme, err := repos.Fs.Read("README.md")
	if err != nil {
		repos.Log.Info("Error: %s", err.Error())
	} else {
		data = readme
	}

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, string(data))
	})

	return e.Start(":3000")
}
