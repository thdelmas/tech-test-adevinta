package config

import (
	"os"
)

// GetServerAddress returns the server address from environment or default
func GetServerAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
