package lsproduct

import (
	"errors"
	"unicode"
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

	return largestProduct(digits, span)
}

func largestProduct(digits string, span int) (int, error) {
	var (
		windowLength  = 0
		windowProduct = 1
		maxProduct    = 0
	)

	for i := 0; i < len(digits); i++ {
		if !unicode.IsNumber(rune(digits[i])) {
			return 0, errors.New("digits must only contain digits")
		}

		digit := parseDigitAt(digits, i)

		// Reset the window
		if digit == 0 {
			windowLength, windowProduct = 0, 1
			continue
		}

		if windowLength < span {
			windowProduct *= digit
			windowLength++
		}

		if windowLength == span {
			if maxProduct < windowProduct {
				maxProduct = windowProduct
			}

			windowProduct /= parseDigitAt(digits, i-span+1)
			windowLength--
		}
	}

	return maxProduct, nil
}

func parseDigitAt(digits string, i int) int {
	return int(digits[i]) - 48
}
