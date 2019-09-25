package isogram

import "unicode"

func IsIsogram(word string) bool {
	var occurrences = make(map[rune]bool)

	for _, r := range word {
		if r == '-' || r == ' ' {
			continue
		}

		r = unicode.ToLower(r)

		if _, ok := occurrences[r]; ok {
			return false
		}

		occurrences[r] = true
	}

	return true
}
