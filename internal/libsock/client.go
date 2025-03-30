package libsock

import "net"

func (e *Sock) Send(text string) error {
	conn, err := net.Dial("unix", e.Path)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.Write([]byte(text)); err != nil {
		return err
	}
	return nil
}
