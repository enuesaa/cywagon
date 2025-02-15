package libfetch

import (
	"fmt"
	"net/http"
)

func Fetch() error {
	res, err := http.Get("https://example.com")
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("faild")
	}
	fmt.Println(res.StatusCode)

	return nil
}
