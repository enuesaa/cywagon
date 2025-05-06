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
		Data: text,
	}
	if err := encoder.Encode(msg); err != nil {
		return err
	}

	var res Message
	if err := decoder.Decode(&res); err != nil {
		return err
	}

	return nil
}
