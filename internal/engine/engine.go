package engine

import (
	"fmt"
	"io"
	"log"
	"net"
)

func RunEngine() error {
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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	bytes, err := io.ReadAll(conn)
	if err != nil {
		log.Panicf("Error: %s", err.Error())
	}
	fmt.Printf("Received: %s", bytes)
}
