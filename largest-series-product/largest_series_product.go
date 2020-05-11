package lsproduct

import (
	"errors"
	"math"
	"strconv"
)

// LargestSeriesProduct calculates the largest product for a
// contiguous substring of digits of length n.
func LargestSeriesProduct(digits string, span int) (int, error) {
	if span < 0 {
		return 0, errors.New("span must be greater than zero")
	}

	if span > len(digits) {
		return 0, errors.New("span must be smaller than the string length")
	}

	if span == 0 {
		return 1, nil
	}

	var product int

	for i := 0; i+span <= len(digits); i++ {
		s := digits[i : i+span]

		n, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}

		if p := productOfDigits(n, span); p > product {
			product = p
		}
	}

	return product, nil
}

func productOfDigits(n, count int) int {
	product := 1

	for count > 0 {
		count--

		x := int(math.Pow10(count))
		digit := n / x
		n -= digit * x

		product *= digit
	}

	return product
}
