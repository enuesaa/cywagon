package handle

import (
	"fmt"
	"io"
	"net"
	"os"
)

var sock = "/var/run/cywagon.sock"


func (h *Handler) Up(paths []string) error {
	confs, err := h.ConfSrv.List(paths)
	if err != nil {
		return err
	}
	h.Engine.PrintBanner(confs)

	if err := h.upCreateSocket(); err != nil {
		return err
	}

	return h.Engine.Serve(confs)
}

func (h *Handler) upCreateSocket() error {
	if _, err := os.Stat(sock); err == nil {
		return fmt.Errorf("sock exists")
	}

	listener, err := net.Listen("unix", sock)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		defer conn.Close()
		buf, err := io.ReadAll(conn)
		if err != nil {
			return err
		}
		h.Log.Info("Received: %s", string(buf))
	}
}
