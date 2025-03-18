package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
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

	// Define expected request
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

	// Mock TrackRequest and GetMostFrequentRequest methods
	mockStatsService.EXPECT().TrackRequest(expectedRequest).Times(1)
	mockStatsService.EXPECT().GetMostFrequentRequest().Return(expectedRequest, 1).Times(1)

	// Initialize StatsHandler with the mock service
	handler := NewStatsHandler(mockStatsService)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/stats", handler.HandleStats)

	// Simulate HTTP request (GET) for stats
	reqBody := `{"int1":3, "int2":5, "limit":15, "str1":"Fizz", "str2":"Buzz"}`
	request, _ := http.NewRequest("GET", "/api/stats", bytes.NewBufferString(reqBody))
	request.Header.Set("Content-Type", "application/json")

	// Capture Response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	// Validate Response using expectedJSON (converted to string)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(expectedJSON), w.Body.String())
}

func TestStatsHandler_InvalidRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock StatsService
	mockStatsService := mocks.NewMockStatsServiceInterface(ctrl)

	// Define various invalid requests with edge cases for Int1 and Int2
	tests := []struct {
		name           string
		reqBody        string
		expectedError  string
		expectedStatus int
		trackTimes     int // Set trackTimes to 0 for error cases and 1 for valid cases
	}{
		{
			name:           "Limit 0",
			reqBody:        `{"int1":3, "int2":5, "limit":0, "str1":"Fizz", "str2":"Buzz"}`,
			expectedError:  "Limit must be greater than 0",
			expectedStatus: http.StatusBadRequest,
			trackTimes:     0, // TrackRequest shouldn't be called
		},
		{
			name:           "Int1 or Int2 zero",
			reqBody:        `{"int1":0, "int2":5, "limit":10, "str1":"Fizz", "str2":"Buzz"}`,
			expectedError:  "Int1 must be greater than 0",
			expectedStatus: http.StatusBadRequest,
			trackTimes:     0, // TrackRequest shouldn't be called
		},
		{
			name:           "Int2 zero",
			reqBody:        `{"int1":3, "int2":0, "limit":10, "str1":"Fizz", "str2":"Buzz"}`,
			expectedError:  "Int2 must be greater than 0",
			expectedStatus: http.StatusBadRequest,
			trackTimes:     0, // TrackRequest shouldn't be called
		},
		{
			name:           "Int1 and Int2 maximum values",
			reqBody:        fmt.Sprintf(`{"int1":%d, "int2":%d, "limit":10, "str1":"Fizz", "str2":"Buzz"}`, math.MaxInt, math.MaxInt),
			expectedError:  "", // Expect valid behavior, so no error
			expectedStatus: http.StatusOK,
			trackTimes:     1, // TrackRequest should be called
		},
		{
			name:           "Limit at max value",
			reqBody:        fmt.Sprintf(`{"int1":3, "int2":5, "limit":%d, "str1":"Fizz", "str2":"Buzz"}`, math.MaxInt),
			expectedError:  "", // Expect valid behavior, so no error
			expectedStatus: http.StatusOK,
			trackTimes:     1, // TrackRequest should be called
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize StatsHandler with the mock service
			handler := NewStatsHandler(mockStatsService)

			// Setup Gin router
			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.GET("/api/stats", handler.HandleStats)

			// Set the expectation based on the trackTimes
			if tt.trackTimes > 0 {
				// For valid cases, mock both TrackRequest and GetMostFrequentRequest
				mockStatsService.EXPECT().TrackRequest(gomock.Any()).Times(tt.trackTimes)
				mockStatsService.EXPECT().GetMostFrequentRequest().Return(
					models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
					0,
				).Times(1)
			} else {
				// Ensure no calls are made to TrackRequest for invalid requests
				mockStatsService.EXPECT().TrackRequest(gomock.Any()).Times(0)
				// No need to mock GetMostFrequentRequest for invalid requests as it shouldn't be called
			}

			// Simulate HTTP request (GET) for stats
			request, _ := http.NewRequest("GET", "/api/stats", bytes.NewBufferString(tt.reqBody))
			request.Header.Set("Content-Type", "application/json")

			// Capture Response
			w := httptest.NewRecorder()
			router.ServeHTTP(w, request)

			// If we expect an error, validate the response
			if tt.expectedError != "" {
				assert.Equal(t, tt.expectedStatus, w.Code)
				assert.JSONEq(t, fmt.Sprintf(`{"error": "%s"}`, tt.expectedError), w.Body.String())
			} else {
				// If no error, validate the successful response (e.g., check for 200 OK)
				assert.Equal(t, tt.expectedStatus, w.Code)
				// Assuming a proper response structure, you can add further validation here if needed
				assert.JSONEq(t, `{"hitCount": 0, "request": {"int1": 3, "int2": 5, "limit": 10, "str1": "Fizz", "str2": "Buzz"}}`, w.Body.String())
			}
		})
	}
}
