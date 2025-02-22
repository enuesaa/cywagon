package libfetch

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func (f *Fetcher) CheckHttpFetch(protocol string, host string, path string) int {
	protocol = strings.ToLower(protocol)
	url := fmt.Sprintf("%s://%s%s", protocol, host, path)

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
