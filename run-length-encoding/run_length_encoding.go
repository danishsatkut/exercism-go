package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode encodes the provided string
func RunLengthEncode(data string) string {
	var encoded strings.Builder

	for len(data) > 0 {
		char := data[0]
		prevLen := len(data)

		data = strings.TrimLeft(data, string(char))
		encoded.WriteString(encodeRun(string(char), prevLen-len(data)))
	}

	return encoded.String()
}

// RunLengthDecode decodes the provided string
func RunLengthDecode(encoded string) string {
	var data strings.Builder
	var runCount int

	for _, r := range encoded {
		if unicode.IsDigit(r) {
			runCount = int(r-'0') + runCount*10
		}

		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			data.WriteString(decodeRun(string(r), runCount))
			runCount = 0
		}
	}

	return data.String()
}

func encodeRun(char string, count int) string {
	var n string

	if count > 1 {
		n = strconv.Itoa(count)
	}

	return n + char
}

func decodeRun(char string, count int) string {
	if count > 0 {
		return strings.Repeat(char, count)
	}

	return char
}
