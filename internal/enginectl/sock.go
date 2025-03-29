package enginectl

import (
	"fmt"
	"io"
	"net"
	"os"
)

var sock = "/tmp/cywagon.sock"

func (e *Engine) StartListenSock() error {
	if _, err := os.Stat(sock); err == nil {
		return fmt.Errorf("sock exists")
	}

	listener, err := net.Listen("unix", sock)
	if err != nil {
		return err
	}
	go e.listenSock(listener)
	return nil
}

func (e *Engine) listenSock(listener net.Listener) {
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			e.Log.Error(err)
			continue
		}
		buf, err := io.ReadAll(conn)
		if err != nil {
			e.Log.Error(err)
			continue
		}
		e.Log.Info("Received: %s", string(buf))
	}
}

func (e *Engine) SendSock() error {
	conn, err := net.Dial("unix", sock)
	if err != nil {
		return err
	}
	defer conn.Close()
	conn.Write([]byte("Hello from client"))

	return nil
}
