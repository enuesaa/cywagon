package handle

import (
	"os"
	"os/signal"
	"syscall"
)

func (h *Handler) Up(paths []string) error {
	confs, err := h.ConfSrv.List(paths)
	if err != nil {
		return err
	}
	h.Engine.PrintBanner(confs)

	if err := h.Engine.StartListenSock(); err != nil {
		return err
	}

	termch := make(chan os.Signal, 1)
	signal.Notify(termch, syscall.SIGINT, syscall.SIGTERM)

	<-termch
	h.Log.Info("close")

	if err := h.Engine.Close(); err != nil {
		return err
	}
	return nil
}
