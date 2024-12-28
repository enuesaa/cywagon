package engine

import (
	"encoding/json"
	"net"
)

type CreateMessage struct {
	Name string `json:"name"`
}

func SendCreate() error {
	message := CreateMessage{
		Name: "hello-this-is-name",
	}
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

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
