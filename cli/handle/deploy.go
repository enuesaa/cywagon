package handle

import (
	"fmt"
	"time"
)

func (h *Handler) Deploy(sitename string, path string) error {
	h.Log.Info("start a new deployment..")

	if err := h.Sock.Send(); err != nil {
		return err
	}
	if err := h.Engine.Deploy(sitename, path); err != nil {
		return err
	}
	time.Sleep(10 * time.Second)

	content, err := h.Engine.Read(sitename)
	if err != nil {
		return err
	}
	fmt.Println(content)

	// ここで socket に命令を送信する

	return nil
}
