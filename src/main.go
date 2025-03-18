package main

import (
	"log"

	"github.com/thdelmas/tech-test-adevinta/config"
	"github.com/thdelmas/tech-test-adevinta/router"
)

func main() {
	log.Println("Adevinta - FizzBuzz")
	router := router.SetupRouter()
	addr := config.GetServerAddress()
	log.Fatal(router.Run(addr))
}
