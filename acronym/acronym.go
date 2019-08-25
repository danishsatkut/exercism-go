// Package acronym implements abbreviation functionality.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate returns abbreviation for a long name.
func Abbreviate(s string) (abbr string) {
	for _, word := range strings.Fields(s) {
		if strings.Contains(word, "-") {
			for _, w := range strings.Split(word, "-") {
				abbr += Abbreviate(w)
			}
		} else if strings.Contains(word, "_") {
			for _, w := range strings.Split(word, "_") {
				abbr += Abbreviate(w)
			}
		} else {
			abbr += string(word[0])
		}
	}

	return strings.ToUpper(abbr)
}
