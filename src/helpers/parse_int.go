package helpers

import (
	"errors"
	"strconv"
)

func ParseInt(value string) (int, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	if intValue <= 0 {
		return 0, errors.New("value must be greater than 0")
	}
	return intValue, nil
}
