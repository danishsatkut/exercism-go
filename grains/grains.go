package grains

import (
	"errors"
	"math"
)

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid square")
	}

	if n == 1 {
		return 1, nil
	}

	return uint64(math.Exp2(float64(n - 1))), nil
}
