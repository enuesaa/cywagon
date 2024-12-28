package msg

import (
	"encoding/json"
	"net"

	"github.com/enuesaa/cywagon/internal/msg/schema"
)

type Sender struct {}

func (s *Sender) send(bytes []byte) error {
	socket, err := Socket()
	if err != nil {
		return err
	}

	conn, err := net.Dial("unix", socket)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sender) SendCreateMessage(name string) error {
	message := schema.Message[schema.CreateData]{
		Operation: "create",
		Data: schema.CreateData{
			Name: name,
		},
	}
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.send(bytes)
}
