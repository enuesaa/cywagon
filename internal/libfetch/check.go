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

func (f *Fetcher) CheckTcpConn(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	return nil
}
