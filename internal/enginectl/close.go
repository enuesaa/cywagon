package enginectl

func (e *Engine) Close() error {
	if e.Sock.Exists() {
		if err := e.Sock.Close(); err != nil {
			return err
		}
	}
	return nil
}
