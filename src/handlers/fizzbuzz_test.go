package handlers

import (
	"bytes"
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

	// Define expected request
	expectedRequest := models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"}
	expectedResponse := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}

	// Mock the GenerateFizzBuzz method for this request
	mockFizzBuzzService.EXPECT().GenerateFizzBuzz(expectedRequest).Return(expectedResponse).Times(1)

	// Initialize FizzBuzzHandler
	handler := NewFizzBuzzHandler(mockFizzBuzzService)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/fizzbuzz", handler.HandleFizzBuzz)

	// Simulate HTTP request (POST with body)
	reqBody := `{"int1":3, "int2":5, "limit":15, "str1":"Fizz", "str2":"Buzz"}`
	request, _ := http.NewRequest("POST", "/api/fizzbuzz", bytes.NewBufferString(reqBody))
	request.Header.Set("Content-Type", "application/json")

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

	// Define invalid request (Limit = 0)
	invalidRequest := models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 0, Str1: "Fizz", Str2: "Buzz"}

	// Setup mock to expect an invalid request and return an empty array (or you can return an error)
	mockFizzBuzzService.EXPECT().GenerateFizzBuzz(invalidRequest).Return([]string{}).Times(1)

	// Initialize FizzBuzzHandler
	handler := NewFizzBuzzHandler(mockFizzBuzzService)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/fizzbuzz", handler.HandleFizzBuzz)

	// Simulate HTTP request (POST with body)
	reqBody := `{"int1":3, "int2":5, "limit":0, "str1":"Fizz", "str2":"Buzz"}`
	request, _ := http.NewRequest("POST", "/api/fizzbuzz", bytes.NewBufferString(reqBody))
	request.Header.Set("Content-Type", "application/json")

	// Capture Response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	// Validate Response: Expecting an empty array because of invalid input (Limit = 0)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `[]`, w.Body.String()) // Return empty array for invalid input
}
