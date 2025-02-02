package ctlconf

import "fmt"

func (c *Conf) RunHandler() error {
	type Response struct {
		Status int `lua:"status"`
	}
	response := Response{
		Status: 404,
	}
	next := func() {
		fmt.Println("this is next function")
	}

	result, err := c.Handler(next, nil, response)
	if err != nil {
		return err
	}
	fmt.Printf("res: %d\n", result.GetInt("status"))

	return nil
}
