package services

import (
	"strconv"
	"sync"

	"github.com/thdelmas/tech-test-adevinta/models"
)

// FizzBuzzServiceInterface defines the methods for FizzBuzzService
type FizzBuzzServiceInterface interface {
	GenerateFizzBuzz(req models.FizzBuzzRequest) []string
}

// FizzBuzzService handles the fizzbuzz business logic
type FizzBuzzService struct{}

// Ensure FizzBuzzService implements the interface
var _ FizzBuzzServiceInterface = (*FizzBuzzService)(nil)

var (
	fizzBuzzServiceInstance *FizzBuzzService
	fizzBuzzServiceOnce     sync.Once
)

// GetFizzBuzzService returns a singleton instance of FizzBuzzService
func GetFizzBuzzService() *FizzBuzzService {
	fizzBuzzServiceOnce.Do(func() {
		fizzBuzzServiceInstance = &FizzBuzzService{}
	})
	return fizzBuzzServiceInstance
}

// GenerateFizzBuzz generates the fizzbuzz sequence based on the given request
func (s *FizzBuzzService) GenerateFizzBuzz(req models.FizzBuzzRequest) []string {
	result := make([]string, req.Limit)
	for i := 1; i <= req.Limit; i++ {
		switch {
		case i%req.Int1 == 0 && i%req.Int2 == 0:
			result[i-1] = req.Str1 + req.Str2
		case i%req.Int1 == 0:
			result[i-1] = req.Str1
		case i%req.Int2 == 0:
			result[i-1] = req.Str2
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}
