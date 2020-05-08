package encode

import (
	"fmt"
	"strconv"
	"unicode"
)

// RunLengthEncode encodes the provided string
func RunLengthEncode(input string) string {
	var encoded string

	var previous rune
	var count = 1

	for i, current := range input {
		if current == previous {
			count++
		} else {
			if i != 0 {
				if count > 1 {
					encoded += strconv.Itoa(count)
				}

				encoded += fmt.Sprintf("%c", previous)
				count = 1
			}
		}

		if i == len(input)-1 {
			if count > 1 {
				encoded += strconv.Itoa(count)
			}

			encoded += fmt.Sprintf("%c", current)
		}

		previous = current
	}

	return encoded
}

// RunLengthDecode decodes the provided string
func RunLengthDecode(input string) string {
	var decoded string
	var previousNumber = 0

	for _, r := range input {
		if unicode.IsDigit(r) {
			previousNumber = int(r-'0') + previousNumber*10
		}

		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			char := string(r)

			if previousNumber > 0 {
				for i := 0; i < previousNumber; i++ {
					decoded += char
				}
			} else {
				decoded += char
			}

			previousNumber = 0
		}
	}

	return decoded
}
