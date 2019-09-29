package isogram

import (
	"unicode"
)

// IsIsogram determines if the word is an isogram.
func IsIsogram(word string) bool {
	var occurrences = make(map[rune]bool)

	for _, r := range word {
		if r == ' ' || r == '-'  {
			continue
		}

		r = unicode.ToLower(r)

		if occurrences[r] {
			return false
		}

		occurrences[r] = true
	}

	return true
}
