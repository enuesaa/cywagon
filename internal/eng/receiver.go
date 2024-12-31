package eng

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/enuesaa/cywagon/internal/msg"
	"github.com/enuesaa/cywagon/internal/repository"
)

type Receiver struct {}

func (r *Receiver) Receive(ctx context.Context, b []byte) error {
	var pre msg.Message[struct{}]
	if err := json.Unmarshal(b, &pre); err != nil {
		return err
	}
	switch pre.Operation {
	case "create":
		return r.ReceiveCreate(ctx, b)
	case "down":
		return r.ReceiveDown(ctx, b)		
	}
	return fmt.Errorf("not found such operation")
}

func (r *Receiver) ReceiveCreate(ctx context.Context, b []byte) error {
	repos := repository.Use(ctx)

	var message msg.Message[msg.CreateData]
	if err := json.Unmarshal(b, &message); err != nil {
		return err
	}
	repos.Log.Info("message: %s", message.Data.Name)

	return nil
}

func (r *Receiver) ReceiveDown(ctx context.Context, b []byte) error {
	if err := Down(ctx); err != nil {
		return err
	}
	return ErrDownEngine
}
