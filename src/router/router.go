package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thdelmas/tech-test-adevinta/handlers"
	"github.com/thdelmas/tech-test-adevinta/services"
)

// SetupRouter configures and returns the Gin router
func SetupRouter(fizzBuzzService services.FizzBuzzServiceInterface, statsService services.StatsServiceInterface) *gin.Engine {
	// Initialize Gin router
	router := gin.New()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	// Define routes
	api := router.Group("/api")
	{
		// Inject dependencies into the handlers
		api.GET("/fizzbuzz", handlers.NewFizzBuzzHandler(fizzBuzzService, statsService).HandleFizzBuzz)
		api.GET("/stats", handlers.NewStatsHandler(statsService).HandleStats)
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
