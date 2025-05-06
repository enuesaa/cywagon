package libsock

import (
	"encoding/json"
	"net"
	"os"
)

type ListenHandler = func(text string) error

func (e *Sock) Listen(handler ListenHandler) error {
	listener, err := net.Listen("unix", e.Path)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			// e.Log.Error(err)
			continue
		}
		if err := e.processConn(conn, handler); err != nil {
			// e.Log.Error(err)
		}
	}
}

func (e *Sock) processConn(conn net.Conn, handler ListenHandler) error {
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	var req Message
	if err := decoder.Decode(&req); err != nil {
		return err
	}
	if err := handler(req.Data); err != nil {
		return err
	}

	res := Message{
		Data: "hello from listener",
	}
	return encoder.Encode(res)
}

func (e *Sock) CloseListener() error {
	return os.Remove(e.Path)
}

func (e *Sock) Exists() bool {
	if _, err := os.Stat(e.Path); err == nil {
		return true
	}
	return false
}
