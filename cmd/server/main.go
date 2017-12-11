package main

import (
	"log"

	"github.com/engineersmy/lookie"
)

// This is a stub for the main process
func main() {
	server := new(lookie.Server)
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
