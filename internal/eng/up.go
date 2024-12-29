package eng

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/enuesaa/cywagon/internal/repository"
)

func Up(ctx context.Context) error {
	repos := repository.Use(ctx)

	pid := os.Getegid()
	if err := repos.Ps.CreatePidFile(pid); err != nil {
		return err
	}

	socket, err := repos.Ps.GetSockPath()
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

	if err := Receive(ctx, bytes); err != nil {
		return err
	}
	return nil
}
