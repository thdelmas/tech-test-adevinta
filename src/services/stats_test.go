package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thdelmas/tech-test-adevinta/models"
)

func TestTrackRequestAndGetMostFrequentRequest(t *testing.T) {
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

func TestGetMostFrequentRequest_NoRequests(t *testing.T) {
	service := NewStatsService()

	// No requests have been tracked yet
	mostFrequent, count, err := service.GetMostFrequentRequest()

	// Verify the error message
	assert.Equal(t, models.FizzBuzzRequest{}, mostFrequent)
	assert.Equal(t, 0, count)
	assert.Error(t, err)

	// Verify the error message
	assert.Equal(t, "no requests have been tracked yet", err.Error())
}
