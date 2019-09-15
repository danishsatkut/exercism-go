// Package raindrops provides functionality to generate raindrops for numbers.
package raindrops

import "strconv"

// Convert converts a number to raindrops string.
func Convert(number int) string {
	var message string

	if isFactorOf(number, 3) {
		message += "Pling"
	}

	if isFactorOf(number, 5) {
		message += "Plang"
	}

	if isFactorOf(number, 7) {
		message += "Plong"
	}

	if message == "" {
		return strconv.Itoa(number)
	}

	return message
}

func isFactorOf(m, n int) bool {
	return m%n == 0
}
