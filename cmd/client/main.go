package main

import (
	"fmt"
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

	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		log.Panicf("Error: %s", err.Error())
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		log.Panicf("Error: %s", err.Error())
	}
}
