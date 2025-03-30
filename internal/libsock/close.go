package libsock

import "os"

func (e *Sock) Close() error {
	return os.Remove(e.Path)
}
