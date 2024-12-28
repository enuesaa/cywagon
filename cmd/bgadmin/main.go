package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Panicf("Error: %s", err.Error())
	}

	socketPath := filepath.Join(homedir, "tmp/cywagon.sock")
	fmt.Printf("%s\n", socketPath)

	listener, err := net.Listen("unix", socketPath)
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
