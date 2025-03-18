package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thdelmas/tech-test-adevinta/services"
)

// StatsHandler struct with service dependency
type StatsHandler struct {
	service services.StatsServiceInterface
}

// NewStatsHandler initializes the handler with a dependency
func NewStatsHandler(service services.StatsServiceInterface) *StatsHandler {
	return &StatsHandler{service: service}
}

func (h *StatsHandler) HandleStats(c *gin.Context) {
	// Call GetMostFrequentRequest method
	mostFrequentRequest, hitCount, err := h.service.GetMostFrequentRequest()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Prepare the response
	response := gin.H{
		"hitCount": hitCount,
		"request":  mostFrequentRequest,
	}

	// Send JSON response
	c.JSON(http.StatusOK, response)
}
