package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thdelmas/tech-test-adevinta/models"
	"github.com/thdelmas/tech-test-adevinta/services"
)

func TestFizzBuzzService_GenerateFizzBuzz(t *testing.T) {
	req := models.FizzBuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 15,
		Str1:  "Fizz",
		Str2:  "Buzz",
	}

	expectedResult := []string{
		"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}

	fizzBuzzService := services.GetFizzBuzzService()
	result := fizzBuzzService.GenerateFizzBuzz(req)

	assert.Equal(t, expectedResult, result)
}
