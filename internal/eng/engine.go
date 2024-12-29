package eng

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/enuesaa/cywagon/internal/msg"
)

func RunEngine(ctx context.Context) error {
	socket, err := Socket()
	if err != nil {
		return err
	}
	listener, err := net.Listen("unix", socket)
	if err != nil {
		log.Panicf("Error: %s", err.Error())
	}
	defer listener.Close()
	fmt.Printf("listening\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panicf("Error: %s", err.Error())
		}

		go handleConnection(ctx, conn)
	}
}

func handleConnection(ctx context.Context, conn net.Conn) error {
	defer conn.Close()

	bytes, err := io.ReadAll(conn)
	if err != nil {
		return err
	}

	receiver := msg.Receiver{}
	if err := receiver.Receive(ctx, bytes); err != nil {
		return err
	}
	return nil
}
