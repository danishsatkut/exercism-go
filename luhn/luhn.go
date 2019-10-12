package luhn

import (
	"strings"
	"unicode"
)

// Valid determines if a number is valid as per the Luhn algorithm.
func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")

	if len(number) <= 1 {
		return false
	}

	return isMod10(number)
}

func isMod10(number string) bool {
	var (
		sum = 0
		pos = len(number) % 2
	)

	for i, r := range number {
		if !unicode.IsNumber(r) {
			return false
		}

		n := int(r - '0')

		// Double every other digit
		if i%2 == pos {
			n = double(n)
		}

		sum += n
	}

	return sum%10 == 0
}

func double(n int) int {
	n *= 2
	if n > 9 {
		n = n - 9
	}

	return n
}
