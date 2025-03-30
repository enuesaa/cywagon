package libsock

import (
	"encoding/json"
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
			e.Log.Error(err)
			continue
		}
		if err := e.handleConn(conn); err != nil {
			e.Log.Error(err)
		}
	}
}

func (e *Sock) handleConn(conn net.Conn) error {
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	var req Message
	if err := decoder.Decode(&req); err != nil {
		return err
	}
	e.Log.Info("listener: ", req.Data)

	res := Message{
		Data: "hello from listener",
	}
	return encoder.Encode(res)
}

func (e *Sock) CloseListener() error {
	return os.Remove(e.Path)
}
