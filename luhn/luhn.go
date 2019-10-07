package luhn

import (
	"strconv"
	"strings"
)

func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")

	if len(number) <= 1 {
		return false
	}

	var digits = make([]int, 0, len(number))

	for i, j := 0, len(number) % 2; i < len(number); i++ {
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

		digits = append(digits, n)
	}

	var sum int

	for _, n := range digits {
		sum += n
	}

	return sum % 10 == 0
}
