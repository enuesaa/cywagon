package msg

import (
	"encoding/json"
	"fmt"

	"github.com/enuesaa/cywagon/internal/msg/schema"
)

func Receive(bytes []byte) error {
	var pre schema.Message[struct{}]
	if err := json.Unmarshal(bytes, &pre); err != nil {
		return err
	}
	if pre.Operation == "create" {
		var message schema.Message[schema.CreateData]
		if err := json.Unmarshal(bytes, &message); err != nil {
			return err
		}
		Log(fmt.Sprintf("message: %s", message.Data.Name))
		return nil
	}

	return fmt.Errorf("not found such operation")
}
