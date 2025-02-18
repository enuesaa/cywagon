package libfetch

import (
	"net"
	"net/http"
)

func (f *Fetcher) CheckHttpFetch(url string) int {
	res, err := http.Get(url)
	if err != nil {
		return 0
	}

	return res.StatusCode
}

func (f *Fetcher) CheckTcpConn() error {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		return err
	}
	defer conn.Close()

	return nil
}
