package utils

import (
	"errors"
	"strconv"
	"strings"
)

func ConvertReceiptsId(str string) (uint, error) {
	parts := strings.Split(str, "-")

	if len(parts) != 2 {
		return 0, errors.New("invalid format: expecting 'x-y'")
	}

	numberStr := parts[1]

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, err
	}

	return uint(number), nil
}
