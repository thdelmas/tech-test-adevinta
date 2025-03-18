package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thdelmas/tech-test-adevinta/models"
)

func TestFizzBuzzService_GenerateFizzBuzz(t *testing.T) {
	tests := []struct {
		name      string
		request   models.FizzBuzzRequest
		expected  []string
		expectErr bool
	}{
		{
			name:      "Standard FizzBuzz",
			request:   models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"},
			expected:  []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
			expectErr: false,
		},
		{
			name:      "Different strings",
			request:   models.FizzBuzzRequest{Int1: 2, Int2: 3, Limit: 10, Str1: "Even", Str2: "Odd"},
			expected:  []string{"1", "Even", "Odd", "Even", "5", "EvenOdd", "7", "Even", "Odd", "Even"},
			expectErr: false,
		},
		{
			name:      "Equal integers",
			request:   models.FizzBuzzRequest{Int1: 2, Int2: 2, Limit: 6, Str1: "X", Str2: "Y"},
			expected:  []string{"1", "XY", "3", "XY", "5", "XY"},
			expectErr: false,
		},
		{
			name:      "Limit 1",
			request:   models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 1, Str1: "Fizz", Str2: "Buzz"},
			expected:  []string{"1"},
			expectErr: false,
		},
		{
			name:    "Large integers",
			request: models.FizzBuzzRequest{Int1: 7, Int2: 11, Limit: 30, Str1: "Seven", Str2: "Eleven"},
			expected: []string{
				"1", "2", "3", "4", "5", "6", "Seven", "8", "9", "10",
				"Eleven", "12", "13", "Seven", "15", "16", "17", "18", "19", "20",
				"Seven", "Eleven", "23", "24", "25", "26", "27", "Seven", "29", "30",
			},
			expectErr: false,
		},
		{
			name:      "Empty strings",
			request:   models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 15, Str1: "", Str2: ""},
			expected:  []string{"1", "2", "", "4", "", "", "7", "8", "", "", "11", "", "13", "14", ""},
			expectErr: false,
		},
		{
			name:      "Zero integers",
			request:   models.FizzBuzzRequest{Int1: 0, Int2: 0, Limit: 15, Str1: "", Str2: ""},
			expected:  nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewFizzBuzzService()
			result, err := service.GenerateFizzBuzz(tt.request)

			if tt.expectErr {
				assert.Error(t, err, "expected an error but got nil")
				assert.Nil(t, result, "expected result to be nil when an error occurs")
			} else {
				assert.NoError(t, err, "did not expect an error but got one")
				assert.Equal(t, tt.expected, result, "unexpected fizzbuzz output")
			}
		})
	}

}
