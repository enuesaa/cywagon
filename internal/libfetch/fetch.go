package libfetch

import (
	"net"
	"net/http"
)

func (f *Fetcher) FetchHTTP(url string) string {
	res, err := http.Get(url)
	if err != nil {
		return "0"
	}

	return res.Status
}

func (f *Fetcher) ConnectTCP(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	return nil
}
