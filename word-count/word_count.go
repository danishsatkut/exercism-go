package wordcount

import (
	"strings"
	"unicode"
)

// Frequency represents count of occurrences of each word
type Frequency map[string]int

// WordCount calculates the occurrence of each word in a phrase
func WordCount(phrase string) Frequency {
	f := Frequency{}

	for _, word := range strings.FieldsFunc(phrase, isSeparator) {
		word = strings.Trim(word, "'")

		f[strings.ToLower(word)]++
	}

	return f
}

func isSeparator(r rune) bool {
	return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
}
