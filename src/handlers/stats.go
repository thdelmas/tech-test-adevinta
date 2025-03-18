package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thdelmas/tech-test-adevinta/models"
	"github.com/thdelmas/tech-test-adevinta/services"
)

// StatsHandler handles the statistics endpoint
func StatsHandler(c *gin.Context) {
	// Get most frequent request
	mostFreqReq, hits := services.GetStatsService().GetMostFrequentRequest()

	// Return response
	c.JSON(http.StatusOK, models.StatsResponse{
		MostFrequentRequest: mostFreqReq,
		Hits:                hits,
	})
}
