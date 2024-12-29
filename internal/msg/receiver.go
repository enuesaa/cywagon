package msg

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/enuesaa/cywagon/internal/msg/schema"
	"github.com/enuesaa/cywagon/internal/repository"
)

type Receiver struct {}

func (r *Receiver) Receive(ctx context.Context, bytes []byte) (string, error) {
	repos := repository.Use(ctx)

	var pre schema.Message[struct{}]
	if err := json.Unmarshal(bytes, &pre); err != nil {
		return "", err
	}
	if pre.Operation == "create" {
		var message schema.Message[schema.CreateData]
		if err := json.Unmarshal(bytes, &message); err != nil {
			return "", err
		}
		repos.Log.Info("message: %s", message.Data.Name)

		return "", nil
	}

	if pre.Operation == "down" {
		return "down", nil
	}

	return "", fmt.Errorf("not found such operation")
}
