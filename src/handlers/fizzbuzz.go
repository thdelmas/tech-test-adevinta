package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thdelmas/tech-test-adevinta/models"
	"github.com/thdelmas/tech-test-adevinta/services"
)

// FizzBuzzHandler struct with service dependency
type FizzBuzzHandler struct {
	service      services.FizzBuzzServiceInterface
	statsService services.StatsServiceInterface
}

// NewFizzBuzzHandler initializes a new handler with dependency injection
func NewFizzBuzzHandler(service services.FizzBuzzServiceInterface, statsService services.StatsServiceInterface) *FizzBuzzHandler {
	return &FizzBuzzHandler{service: service, statsService: statsService}
}

func (h *FizzBuzzHandler) HandleFizzBuzz(c *gin.Context) {
	var req models.FizzBuzzRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate that Limit is greater than 0
	if req.Limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	h.statsService.TrackRequest(req)

	result, err := h.service.GenerateFizzBuzz(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
