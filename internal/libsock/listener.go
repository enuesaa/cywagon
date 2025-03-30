package libsock

import (
	"io"
	"net"
	"os"
)

func (e *Sock) Listen() error {
	listener, err := net.Listen("unix", e.Path)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		buf, err := io.ReadAll(conn)
		if err != nil {
			return err
		}
		e.Log.Info("Received: %s", string(buf))
	}
}

func (e *Sock) CloseListener() error {
	return os.Remove(e.Path)
}
