package libsock

import "os"

func (e *Sock) Exists() bool {
	if _, err := os.Stat(e.Path); err == nil {
		return true
	}
	return false
}
