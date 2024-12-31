package msg

import (
	"context"
	"encoding/json"
	"net"

	"github.com/enuesaa/cywagon/internal/msg/schema"
	"github.com/enuesaa/cywagon/internal/repository"
)

type Sender struct {}

func (s *Sender) send(ctx context.Context, bytes []byte) error {
	repos := repository.Use(ctx)

	socket, err := repos.Ps.GetSockPath()
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

func (s *Sender) SendCreateMessage(ctx context.Context, name string) error {
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
	return s.send(ctx, bytes)
}
