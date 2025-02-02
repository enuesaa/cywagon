package ctlconf

func RunHandler(conf Conf) error {
	type Response struct {
		Status int `lua:"status"`
	}
	// response := Response{
	// 	Status: 404,
	// }
	// next := func() {
	// 	fmt.Println("this is next function")
	// }

	// result, err := conf.Handler.Run(next, nil, response)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("res: %d\n", result.GetInt("status"))

	return nil
}
