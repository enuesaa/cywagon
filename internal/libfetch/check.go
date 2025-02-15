package libfetch

import (
	"fmt"
	"net"
	"net/http"
)

func (f *Fetcher) CheckHttpFetch() error {
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

func (f *Fetcher) CheckTcpConn() error {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		return err
	}
	defer conn.Close()

	return nil
}
