package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello from client")

	client_addr := os.Getenv("CLIENT_ORIGIN")
	if client_addr == "" {
		log.Fatal("$CLIENT_ORIGIN is not set")
	}

	log.Printf("client address: %s", client_addr)
}
