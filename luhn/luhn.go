package luhn

import (
	"strconv"
	"strings"
)

// Valid determines if a number is valid as per the Luhn algorithm.
func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")

	var totalDigits = len(number)

	if totalDigits <= 1 {
		return false
	}

	var sum int

	for i, j := 0, totalDigits % 2; i < totalDigits; i++ {
		n, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return false
		}

		if i == j {
			n = n * 2
			if n > 9 {
				n = n - 9
			}

			j += 2
		}

		sum += n
	}

	return sum % 10 == 0
}
