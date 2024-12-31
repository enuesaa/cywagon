package engctl

import (
	"context"
	"encoding/json"

	"github.com/enuesaa/cywagon/internal/msg"
	"github.com/enuesaa/cywagon/internal/repository"
)

func SendDownMessage(ctx context.Context) error {
	repos := repository.Use(ctx)

	message := msg.Message[msg.DownData]{
		Operation: "down",
		Data: msg.DownData{},
	}
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return repos.Ps.SendThroughSocket(bytes)
}
