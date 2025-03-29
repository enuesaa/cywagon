package enginectl

import "fmt"

func (e *Engine) StartListenSock() error {
	if e.Sock.Exists() {
		return fmt.Errorf("sock exists")
	}
	go e.Sock.Listen()

	return nil
}
