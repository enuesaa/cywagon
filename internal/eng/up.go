package eng

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/enuesaa/cywagon/internal/msg"
)

func Up(ctx context.Context) error {
	pid := os.Getegid()
	if err := CreatePidFile(pid); err != nil {
		return err
	}

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
	operation, err := receiver.Receive(ctx, bytes)
	if err != nil {
		return err
	}
	if operation == "down" {
		if err := Down(); err != nil {
			return err
		}
		os.Exit(0)
	}
	return nil
}
