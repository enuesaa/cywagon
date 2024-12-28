package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/cywagon/internal"
)

func Run() {
	if len(os.Args) == 1 {
		log.Fatalf("Error: missing required arg\n")
	}
	cmd := os.Args[1]

	switch cmd {
	case "engine-start":
		if err := internal.RunEngine(); err != nil {
			log.Fatalf("Error: %s\n", err.Error())
		}
	case "up":
		if err := internal.Up(); err != nil {
			log.Fatalf("Error: %s\n", err.Error())
		}
	case "hello":
		if err := internal.Hello(); err != nil {
			log.Fatalf("Error: %s\n", err.Error())
		}
	}
	fmt.Println(os.Args)
}
