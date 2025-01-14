package conf

import (
	"fmt"

	"github.com/enuesaa/cywagon/internal/liblua"
)

type Config struct {
	Hostname string
	Port int
	handler liblua.Fn
}

func (c *Config) RunHandler() error {
	type Response struct {
		Status int `lua:"status"`
	}
	response := Response{
		Status: 404,
	}

	result, err := c.handler.Run(Next, nil, response)
	if err != nil {
		return err
	}
	fmt.Printf("res: %d\n", result.GetInt("status"))

	return nil
}

func Next() {
	fmt.Println("this is next function")
}