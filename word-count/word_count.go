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

	words := strings.FieldsFunc(phrase, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})

	for _, word := range words {
		word = strings.Trim(word, "'")

		f[strings.ToLower(word)]++
	}

	return f
}
