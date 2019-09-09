package raindrops

import "fmt"

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
		return fmt.Sprintf("%d", number)
	}

	return message
}

func isFactorOf(m, n int) bool {
	return m%n == 0
}
