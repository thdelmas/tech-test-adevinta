package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thdelmas/tech-test-adevinta/models"
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
	var request models.FizzBuzzRequest

	// Bind request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate request
	if request.Int1 <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Int1 must be greater than 0"})
		return // Ensure you return here before calling TrackRequest
	}

	if request.Int2 <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Int2 must be greater than 0"})
		return // Ensure you return here before calling TrackRequest
	}

	if request.Limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Limit must be greater than 0"})
		return // Ensure you return here before calling TrackRequest
	}

	// Only call TrackRequest if validation passes
	h.service.TrackRequest(request)

	// Rest of your handler logic
	mostFreqReq, hitCount := h.service.GetMostFrequentRequest()

	c.JSON(http.StatusOK, gin.H{
		"hitCount": hitCount,
		"request":  mostFreqReq,
	})
}
