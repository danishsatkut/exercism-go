// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	r := strings.Fields(s)

	//log.Printf("Fields: %v, Count: %v, First: %v", r, len(r), string(r[0][0]))

	var abbr string

	for _, word := range r {
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
