// Package acronym implements abbreviation functionality.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

const separators = " -_"

// Abbreviate returns abbreviation for a long name.
func Abbreviate(s string) (abbr string) {
	for _, word := range strings.FieldsFunc(s, isSeparator) {
		abbr += string(word[0])
	}

	return strings.ToUpper(abbr)
}

func isSeparator(r rune) bool {
	return strings.ContainsRune(separators, r)
}
