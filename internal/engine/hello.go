package engine

import (
	"net"
)

func Hello() error {
	socket, err := Socket()
	if err != nil {
		return err
	}

	conn, err := net.Dial("unix", socket)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		return err
	}
	return nil
}
