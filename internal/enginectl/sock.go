package enginectl

import "fmt"

func (e *Engine) StartListenSock() error {
	if e.Sock.Exists() {
		return fmt.Errorf("sock exists")
	}
	go e.Sock.Listen(func(text string) error {
		e.Log.Info("text: ", text)
		return nil
	})

	return nil
}
