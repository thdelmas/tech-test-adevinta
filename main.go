package main

import (
	"log"

	"github.com/thdelmas/tech-test-adevinta/config"
	"github.com/thdelmas/tech-test-adevinta/router"
)

func main() {
	// Initialize router with Gin
	router := router.SetupRouter()

	// Start server
	addr := config.GetServerAddress()
	log.Printf("Starting FizzBuzz server on %s", addr)
	log.Fatal(router.Run(addr))
}
