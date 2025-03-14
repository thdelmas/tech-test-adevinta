package helpers

import "strconv"

// Helper function to parse integer parameters
func ParseInt(value string) (int, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil || intValue <= 0 {
		return 0, err
	}
	return intValue, nil
}
