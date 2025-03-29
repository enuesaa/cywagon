package libsock

import (
	"io"
	"net"
	"os"

	"github.com/enuesaa/cywagon/internal/infra"
)

func New() Sock {
	return Sock{
		Container: infra.Default,
		Path: "/tmp/cywagon.sock",
	}
}

type Sock struct {
	infra.Container

	Path string
}

func (e *Sock) Exists() bool {
	if _, err := os.Stat(e.Path); err == nil {
		return true
	}
	return false
}

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

func (e *Sock) Send() error {
	conn, err := net.Dial("unix", e.Path)
	if err != nil {
		return err
	}
	defer conn.Close()
	conn.Write([]byte("Hello from client"))

	return nil
}
