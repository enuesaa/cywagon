package eng

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/enuesaa/cywagon/internal/msg"
	"github.com/enuesaa/cywagon/internal/repository"
)

func Receive(ctx context.Context, bytes []byte) error {
	repos := repository.Use(ctx)

	var pre msg.Message[struct{}]
	if err := json.Unmarshal(bytes, &pre); err != nil {
		return err
	}
	if pre.Operation == "create" {
		var message msg.Message[msg.CreateData]
		if err := json.Unmarshal(bytes, &message); err != nil {
			return err
		}
		repos.Log.Info("message: %s", message.Data.Name)

		return nil
	}

	if pre.Operation == "down" {
		if err := Down(ctx); err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("not found such operation")
}
