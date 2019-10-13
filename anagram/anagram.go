package anagram

import (
	"strings"
	"unicode"
)

// Detect returns anagrams for a word from possible candidates.
func Detect(word string, candidates []string) []string {
	var anagrams = make([]string, 0, len(candidates))

	word = strings.ToUpper(word)
	s := newLetterCount(word)

	for i, candidate := range candidates {
		if len(word) != len(candidate) {
			continue
		}

		candidate = strings.ToUpper(candidate)

		if word == candidate {
			continue
		}

		if s == newLetterCount(candidate) {
			anagrams = append(anagrams, candidates[i])
		}
	}

	return anagrams
}

type letterCount [26]int

func newLetterCount(word string) letterCount {
	var letters = letterCount{}

	for _, r := range word {
		if unicode.IsLetter(r) {
			letters[r - 'A']++
		}
	}

	return letters
}
