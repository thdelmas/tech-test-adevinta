package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/thdelmas/tech-test-adevinta/mocks"
	"github.com/thdelmas/tech-test-adevinta/models"
)

// TestStatsHandler_Success tests the StatsHandler for a successful request
func TestStatsHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock StatsService
	mockStatsService := mocks.NewMockStatsServiceInterface(ctrl)

	// Define expected request data that GetMostFrequentRequest should return
	expectedRequest := models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"}

	// Define expected response
	expectedResponse := gin.H{
		"hitCount": 1,
		"request":  expectedRequest,
	}

	// Convert expectedResponse to JSON string
	expectedJSON, err := json.Marshal(expectedResponse)
	if err != nil {
		t.Fatalf("Error marshalling expected response: %v", err)
	}

	// Mock GetMostFrequentRequest method to return the expected data
	mockStatsService.EXPECT().GetMostFrequentRequest().Return(expectedRequest, 1, nil).Times(1)

	// Initialize StatsHandler with the mock service
	handler := NewStatsHandler(mockStatsService)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/stats", handler.HandleStats)

	// Simulate HTTP request (GET) for stats
	request, _ := http.NewRequest("GET", "/api/stats", nil)

	// Capture Response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	// Validate Response using expectedJSON (converted to string)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(expectedJSON), w.Body.String())

	// Ensure TrackRequest was NOT called by verifying no unexpected calls were made
	mockStatsService.EXPECT().TrackRequest(gomock.Any()).Times(0) // Ensure it's never called
}
