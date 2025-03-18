package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thdelmas/tech-test-adevinta/models"
	"github.com/thdelmas/tech-test-adevinta/services"
)

// FizzBuzzHandler struct with service dependency
type FizzBuzzHandler struct {
	service services.FizzBuzzServiceInterface
}

// NewFizzBuzzHandler initializes a new handler with dependency injection
func NewFizzBuzzHandler(service services.FizzBuzzServiceInterface) *FizzBuzzHandler {
	return &FizzBuzzHandler{service: service}
}

// HandleFizzBuzz processes requests
func (h *FizzBuzzHandler) HandleFizzBuzz(c *gin.Context) {
	var req models.FizzBuzzRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result := h.service.GenerateFizzBuzz(req)
	c.JSON(http.StatusOK, result)
}
