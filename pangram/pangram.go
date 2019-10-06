package pangram

import "unicode"

func IsPangram(sentence string) bool {
	if sentence == "" {
		return false
	}

	var occurence = make(map[rune]bool)

	for _, l := range sentence {
		if unicode.IsSpace(l) || l == '_' || unicode.IsDigit(l) || unicode.IsPunct(l) {
			continue
		}

		occurence[unicode.ToLower(l)] = true
	}

	return len(occurence) == 26
}
