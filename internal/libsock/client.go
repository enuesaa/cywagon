package libsock

import (
	"encoding/json"
	"net"
)

func (e *Sock) Send(text string) error {
	conn, err := net.Dial("unix", e.Path)
	if err != nil {
		return err
	}
	defer conn.Close()
	
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	msg := Message{
		Data: "hello from client",
	}
	if err := encoder.Encode(msg); err != nil {
		return err
	}
	
	var res Message
	if err := decoder.Decode(&res); err != nil {
		return err
	}
	e.Log.Info("client: ", res.Data)

	return nil
}
