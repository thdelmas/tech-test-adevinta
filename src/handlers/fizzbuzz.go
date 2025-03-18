package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thdelmas/tech-test-adevinta/helpers"
	"github.com/thdelmas/tech-test-adevinta/models"
	"github.com/thdelmas/tech-test-adevinta/services"
)

// FizzBuzzHandler handles the fizzbuzz endpoint
func FizzBuzzHandler(c *gin.Context) {
	// Parse and validate query parameters
	var request models.FizzBuzzRequest

	// Bind and validate int1
	int1, err := helpers.ParseInt(c.Query("int1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid int1 parameter"})
		return
	}
	request.Int1 = int1

	// Bind and validate int2
	int2, err := helpers.ParseInt(c.Query("int2"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid int2 parameter"})
		return
	}
	request.Int2 = int2

	// Bind and validate limit
	limit, err := helpers.ParseInt(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}
	request.Limit = limit

	// Bind and validate str1
	str1 := c.Query("str1")
	if str1 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing str1 parameter"})
		return
	}
	request.Str1 = str1

	// Bind and validate str2
	str2 := c.Query("str2")
	if str2 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing str2 parameter"})
		return
	}
	request.Str2 = str2

	// Track this request
	services.GetStatsService().TrackRequest(request)

	// Generate the fizzbuzz sequence
	result := services.GetFizzBuzzService().GenerateFizzBuzz(request)

	// Return response
	c.JSON(http.StatusOK, models.FizzBuzzResponse{
		Result: result,
	})
}
