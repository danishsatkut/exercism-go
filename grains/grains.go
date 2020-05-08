package grains

import (
	"errors"
	"math"
)

// Square calculates the number of grains for a given square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid square")
	}

	return uint64(math.Exp2(float64(n - 1))), nil
}

// Total calculates the total number of grains for the entire chessboard
func Total() uint64 {
	var total uint64

	for i := 1; i <= 64; i++ {
		if value, err := Square(i); err == nil {
			total += value
		}
	}

	return total
}
