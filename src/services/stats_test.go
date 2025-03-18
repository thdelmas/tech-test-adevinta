package services

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thdelmas/tech-test-adevinta/models"
)

func TestTrackRequest_Success(t *testing.T) {
	service := NewStatsService()

	// Define test requests
	req1 := models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"}
	req2 := models.FizzBuzzRequest{Int1: 2, Int2: 4, Limit: 8, Str1: "Foo", Str2: "Bar"}

	// Track requests
	service.TrackRequest(req1)
	service.TrackRequest(req1)
	service.TrackRequest(req2)

	// Get most frequent request
	mostFrequent, count, err := service.GetMostFrequentRequest()

	// Verify that there is no error
	assert.NoError(t, err)

	// Verify the most frequent request is req1
	assert.Equal(t, req1, mostFrequent)
	assert.Equal(t, 2, count)
}
func TestGetMostFrequentRequest_Errors(t *testing.T) {
	tests := []struct {
		name          string
		trackRequests func(service *StatsService)
		expectedMost  models.FizzBuzzRequest
		expectedCount int
		expectedErr   string
	}{
		{
			name:          "No requests tracked",
			trackRequests: func(service *StatsService) {},
			expectedMost:  models.FizzBuzzRequest{},
			expectedCount: 0,
			expectedErr:   "no requests have been tracked yet",
		},
		{
			name: "Track empty request",
			trackRequests: func(service *StatsService) {
				service.TrackRequest(models.FizzBuzzRequest{})
			},
			expectedMost:  models.FizzBuzzRequest{},
			expectedCount: 1,
			expectedErr:   "",
		},
		{
			name: "Track multiple requests with equal frequency",
			trackRequests: func(service *StatsService) {
				req1 := models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"}
				req2 := models.FizzBuzzRequest{Int1: 2, Int2: 4, Limit: 8, Str1: "Foo", Str2: "Bar"}
				service.TrackRequest(req1)
				service.TrackRequest(req2)
				service.TrackRequest(req1)
				service.TrackRequest(req2)
			},
			expectedMost:  models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
			expectedCount: 2,
			expectedErr:   "",
		},
		{
			name: "Track requests concurrently",
			trackRequests: func(service *StatsService) {
				req := models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"}
				var wg sync.WaitGroup
				for i := 0; i < 1000; i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						service.TrackRequest(req)
					}()
				}
				wg.Wait()
			},
			expectedMost:  models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
			expectedCount: 1000,
			expectedErr:   "",
		},
		{
			name: "Track identical requests and check updated count",
			trackRequests: func(service *StatsService) {
				req := models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"}
				service.TrackRequest(req)
				service.TrackRequest(req)
			},
			expectedMost:  models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
			expectedCount: 2,
			expectedErr:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewStatsService()
			tt.trackRequests(service)

			mostFrequent, count, err := service.GetMostFrequentRequest()

			if tt.expectedErr != "" {
				assert.Error(t, err, "expected an error but got nil")
				assert.Contains(t, err.Error(), tt.expectedErr, "unexpected error message")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedMost, mostFrequent)
				assert.Equal(t, tt.expectedCount, count)
			}
		})
	}
}
