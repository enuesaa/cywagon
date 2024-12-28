package msg

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/enuesaa/cywagon/internal/msg/schema"
	"github.com/enuesaa/cywagon/internal/repository"
)

type Receiver struct {}

func (r *Receiver) Receive(ctx context.Context, bytes []byte) error {
	var pre schema.Message[struct{}]
	if err := json.Unmarshal(bytes, &pre); err != nil {
		return err
	}
	if pre.Operation == "create" {
		var message schema.Message[schema.CreateData]
		if err := json.Unmarshal(bytes, &message); err != nil {
			return err
		}
		logrepo := repository.UseLog(ctx)
		logrepo.Info("message: %s", message.Data.Name)

		return nil
	}

	return fmt.Errorf("not found such operation")
}
