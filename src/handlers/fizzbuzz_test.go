package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/thdelmas/tech-test-adevinta/mocks"
	"github.com/thdelmas/tech-test-adevinta/models"
)

func TestFizzBuzzHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock FizzBuzzService
	mockFizzBuzzService := mocks.NewMockFizzBuzzServiceInterface(ctrl)
	mockStatsService := mocks.NewMockStatsServiceInterface(ctrl)

	// Define expected request
	expectedRequest := models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"}
	expectedResponse := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}

	// Mock the GenerateFizzBuzz method for this request
	mockFizzBuzzService.EXPECT().GenerateFizzBuzz(expectedRequest).Return(expectedResponse, nil).Times(1)
	mockStatsService.EXPECT().TrackRequest(expectedRequest).Times(1)

	// Initialize FizzBuzzHandler
	handler := NewFizzBuzzHandler(mockFizzBuzzService, mockStatsService)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/fizzbuzz", handler.HandleFizzBuzz) // Use GET for query parameters

	// Simulate HTTP request (GET with query parameters)
	request, _ := http.NewRequest("GET", "/api/fizzbuzz?int1=3&int2=5&limit=15&str1=Fizz&str2=Buzz", nil)

	// Capture Response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	// Validate Response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"]`, w.Body.String())
}

func TestFizzBuzzHandler_InvalidRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock FizzBuzzService
	mockFizzBuzzService := mocks.NewMockFizzBuzzServiceInterface(ctrl)
	mockStatsService := mocks.NewMockStatsServiceInterface(ctrl)

	// Define invalid request: Limit is set to 0 to simulate an invalid input
	invalidRequest := models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 0, Str1: "Fizz", Str2: "Buzz"}

	// Expect the GenerateFizzBuzz method to not be called in case of an invalid request
	mockFizzBuzzService.EXPECT().GenerateFizzBuzz(invalidRequest).Times(0)

	// Initialize FizzBuzzHandler
	handler := NewFizzBuzzHandler(mockFizzBuzzService, mockStatsService)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/fizzbuzz", handler.HandleFizzBuzz) // Use GET for query parameters

	// Simulate HTTP request (GET with invalid query parameters)
	request, _ := http.NewRequest("GET", "/api/fizzbuzz?int1=3&int2=5&limit=0&str1=Fizz&str2=Buzz", nil)

	// Capture Response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	// Validate Response
	assert.Equal(t, http.StatusBadRequest, w.Code) // Expecting BadRequest due to invalid limit
	assert.JSONEq(t, `{"error": "Invalid request"}`, w.Body.String())
}
