package luhn

import (
	"strings"
	"unicode"
)

// Valid determines if a number is valid as per the Luhn algorithm.
func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")

	var totalDigits = len(number)
	if totalDigits <= 1 {
		return false
	}

	var sum int
	var pos = totalDigits % 2

	for i, r := range number {
		if !unicode.IsNumber(r) {
			return false
		}

		n := int(r - '0')

		if i == pos {
			n = n * 2
			if n > 9 {
				n = n - 9
			}

			pos += 2
		}

		sum += n
	}

	return sum % 10 == 0
}
