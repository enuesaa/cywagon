package msg

import (
	"context"
	"encoding/json"

	"github.com/enuesaa/cywagon/internal/msg/schema"
	"github.com/enuesaa/cywagon/internal/repository"
)

type Sender struct {}

func (s *Sender) SendCreateMessage(ctx context.Context, name string) error {
	repos := repository.Use(ctx)

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
	return repos.Ps.SendThroughSocket(bytes)
}
