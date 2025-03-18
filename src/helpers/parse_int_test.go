package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedValue int
		expectError   bool
	}{
		{
			name:          "Valid positive integer",
			input:         "42",
			expectedValue: 42,
			expectError:   false,
		},
		{
			name:          "Zero value",
			input:         "0",
			expectedValue: 0,
			expectError:   true, // Function should return error for <= 0
		},
		{
			name:          "Negative integer",
			input:         "-5",
			expectedValue: 0,
			expectError:   true, // Function should return error for <= 0
		},
		{
			name:          "Non-integer string",
			input:         "abc",
			expectedValue: 0,
			expectError:   true,
		},
		{
			name:          "Float string",
			input:         "3.14",
			expectedValue: 0,
			expectError:   true,
		},
		{
			name:          "Empty string",
			input:         "",
			expectedValue: 0,
			expectError:   true,
		},
		{
			name:          "Very large integer",
			input:         "9223372036854775807", // Max int64
			expectedValue: 9223372036854775807,   // This may be platform dependent (int vs int64)
			expectError:   false,
		},
		{
			name:          "Integer with leading/trailing spaces",
			input:         " 123 ",
			expectedValue: 0,
			expectError:   true, // strconv.Atoi doesn't trim spaces
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := ParseInt(tt.input)

			if tt.expectError {
				assert.Error(t, err, "Expected an error for input: %s", tt.input)
				assert.Equal(t, 0, value, "Value should be 0 when error is returned")
			} else {
				assert.NoError(t, err, "Did not expect an error for input: %s", tt.input)
				assert.Equal(t, tt.expectedValue, value, "Value does not match expected for input: %s", tt.input)
			}
		})
	}
}
