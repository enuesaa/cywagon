package libsock

import "net"

func (e *Sock) Send() error {
	conn, err := net.Dial("unix", e.Path)
	if err != nil {
		return err
	}
	defer conn.Close()
	conn.Write([]byte("Hello from client"))

	return nil
}
