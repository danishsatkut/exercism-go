package raindrops

import "fmt"

func Convert(n int) string {
	var message string

	if n%3 == 0 {
		message += "Pling"
	}

	if n%5 == 0 {
		message += "Plang"
	}

	if n%7 == 0 {
		message += "Plong"
	}

	if message == "" {
		message = fmt.Sprintf("%d", n)
	}

	return message
}
