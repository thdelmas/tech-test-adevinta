package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thdelmas/tech-test-adevinta/handlers"
)

// SetupRouter configures and returns the Gin router
func SetupRouter() *gin.Engine {
	// Set Gin to release mode in production
	// gin.SetMode(gin.ReleaseMode)

	// Initialize Gin router
	router := gin.New()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	// Define routes
	api := router.Group("/api")
	{
		api.GET("/fizzbuzz", handlers.FizzBuzzHandler)
		api.GET("/stats", handlers.StatsHandler)
	}

	return router
}

// CORSMiddleware adds CORS headers to responses
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
