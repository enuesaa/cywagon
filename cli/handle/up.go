package handle

import "time"

func (h *Handler) Up(paths []string) error {
	confs, err := h.ConfSrv.List(paths)
	if err != nil {
		return err
	}
	h.Engine.PrintBanner(confs)

	if err := h.Engine.StartListenSock(); err != nil {
		return err
	}
	// return h.Engine.Serve(confs)

	time.Sleep(100*time.Second)
	
	return nil
}
