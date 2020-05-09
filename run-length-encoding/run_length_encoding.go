package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode encodes the provided string
func RunLengthEncode(input string) string {
	var encoded strings.Builder

	var previous rune
	var count = 1

	for i, current := range input {
		if current == previous {
			count++
		} else {
			if i != 0 {
				encoded.WriteString(encodeRun(previous, count))
				count = 1
			}
		}

		if i == len(input)-1 {
			encoded.WriteString(encodeRun(current, count))
		}

		previous = current
	}

	return encoded.String()
}

// RunLengthDecode decodes the provided string
func RunLengthDecode(input string) string {
	var decoded strings.Builder
	var runCount = 0

	for _, r := range input {
		if unicode.IsDigit(r) {
			runCount = int(r-'0') + runCount*10
		}

		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			decoded.WriteString(decodeRun(string(r), runCount))
			runCount = 0
		}
	}

	return decoded.String()
}

func encodeRun(char rune, count int) string {
	var e string

	if count > 1 {
		e += strconv.Itoa(count)
	}

	return e + fmt.Sprintf("%c", char)
}

func decodeRun(char string, count int) string {
	if count > 0 {
		return strings.Repeat(char, count)
	}

	return char
}
