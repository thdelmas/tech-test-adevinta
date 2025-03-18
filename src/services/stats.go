package services

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/thdelmas/tech-test-adevinta/models"
)

// StatsServiceInterface defines methods for tracking and retrieving statistics
type StatsServiceInterface interface {
	TrackRequest(req models.FizzBuzzRequest)
	GetMostFrequentRequest() (models.FizzBuzzRequest, int)
}

// StatsService tracks and provides statistics for requests
type StatsService struct {
	counter map[string]int
	mutex   sync.RWMutex
}

// Ensure StatsService implements the interface
var _ StatsServiceInterface = (*StatsService)(nil)

// NewStatsService creates a new instance of StatsService (avoids using a singleton)
func NewStatsService() *StatsService {
	return &StatsService{
		counter: make(map[string]int),
	}
}

// TrackRequest tracks a request in the statistics
func (s *StatsService) TrackRequest(req models.FizzBuzzRequest) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	key := fmt.Sprintf("%d:%d:%d:%s:%s", req.Int1, req.Int2, req.Limit, req.Str1, req.Str2)
	s.counter[key]++
}

// GetMostFrequentRequest returns the most frequent request and its hit count
func (s *StatsService) GetMostFrequentRequest() (models.FizzBuzzRequest, int) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var maxKey string
	maxHits := 0

	for key, hits := range s.counter {
		if hits > maxHits {
			maxKey = key
			maxHits = hits
		}
	}

	if maxHits == 0 {
		return models.FizzBuzzRequest{}, 0
	}

	var req models.FizzBuzzRequest
	parts := strings.Split(maxKey, ":")
	if len(parts) == 5 {
		req.Int1, _ = strconv.Atoi(parts[0])
		req.Int2, _ = strconv.Atoi(parts[1])
		req.Limit, _ = strconv.Atoi(parts[2])
		req.Str1 = parts[3]
		req.Str2 = parts[4]
	}

	return req, maxHits
}
