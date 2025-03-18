package main

import (
	"log"

	"github.com/thdelmas/tech-test-adevinta/config"
	"github.com/thdelmas/tech-test-adevinta/router"
	"github.com/thdelmas/tech-test-adevinta/services"
)

func main() {
	log.Println("Adevinta - FizzBuzz")
	fizzBuzzService := services.NewFizzBuzzService()
	statsService := services.NewStatsService()

	// Set up the router with dependencies injected
	r := router.SetupRouter(fizzBuzzService, statsService)
	addr := config.GetServerAddress()
	log.Fatal(r.Run(addr))
}
