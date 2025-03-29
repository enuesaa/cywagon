package handle

import (
	"fmt"
	"net"
	"time"
)

func (h *Handler) Deploy(sitename string, path string) error {
	h.Log.Info("start a new deployment..")

	conn, err := net.Dial("unix", sock)
	if err != nil {
		return err
	}
	defer conn.Close()
	conn.Write([]byte("Hello from client"))

	time.Sleep(100 * time.Second)


	//  client としてデーモンに指示する形式かな

	if err := h.Engine.Deploy(sitename, path); err != nil {
		return err
	}
	time.Sleep(10 * time.Second)

	content, err := h.Engine.Read(sitename)
	if err != nil {
		return err
	}
	fmt.Println(content)

	return nil
}
