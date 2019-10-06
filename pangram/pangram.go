package pangram

import "unicode"

// IsPangram determines if a sentence is a pangram.
// A pangram is a sentence using every letter of the alphabet at least once.
func IsPangram(sentence string) bool {
	var alphabets [26]int

	for _, l := range sentence {
		if unicode.IsLetter(l) {
			alphabets[unicode.ToLower(l) - 'a']++
		}
	}

	for _, value := range alphabets {
		if value == 0 {
			return false
		}
	}

	return true
}
