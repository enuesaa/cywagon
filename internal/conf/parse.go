package conf

import (
	"context"
	"fmt"

	"github.com/enuesaa/cywagon/internal/liblua"
	"github.com/enuesaa/cywagon/internal/repository"
)

func Parse(ctx context.Context) error {
	repos := repository.Use(ctx)
	scriptbytes, err := repos.Fs.Read("testdata/sites-enabled/example.lua")
	if err != nil {
		return err
	}

	runner := liblua.NewRunner(string(scriptbytes))
	if err := runner.Run(); err != nil {
		return err
	}

	fmt.Printf("hostname: %s\n", runner.GetString("hostname"))
	fmt.Printf("port: %d\n", runner.GetInt("port"))

	type Response struct {
		Status int `lua:"status"`
	}
	response := Response{
		Status: 404,
	}
	result, err := runner.RunFunction("handle", Next, nil, response)
	if err != nil {
		return err
	}
	fmt.Printf("res: %d\n", result.GetInt("status"))

	return nil
}

type Config struct {}

func Next() {
	fmt.Println("this is next function")
}
