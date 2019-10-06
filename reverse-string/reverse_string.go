package reverse

import "bytes"

// Reverse performs string reversal
func Reverse(s string) string {
	var output bytes.Buffer

	r := []rune(s)

	for i := len(r) - 1; i >= 0; i-- {
		output.WriteRune(r[i])
	}

	return output.String()
}
