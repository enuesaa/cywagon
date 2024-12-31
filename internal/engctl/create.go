package engctl

import (
	"context"
	"encoding/json"

	"github.com/enuesaa/cywagon/internal/msg"
	"github.com/enuesaa/cywagon/internal/repository"
)

func SendCreateMessage(ctx context.Context) error {
	repos := repository.Use(ctx)

	message := msg.Message[msg.CreateData]{
		Operation: "create",
		Data: msg.CreateData{
			Name: "aaa",
		},
	}
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return repos.Ps.SendThroughSocket(bytes)
}
