package anagram

import (
	"strings"
	"unicode"
)

// Detect returns anagrams for a word from possible candidates.
func Detect(word string, candidates []string) []string {
	var anagrams = make([]string, 0, len(candidates))

	sc := newLetterCount(word)

	for _, candidate := range candidates {
		if len(word) != len(candidate) || strings.EqualFold(word, candidate) {
			continue
		}

		if sc == newLetterCount(candidate) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

type letterCount [26]int

func newLetterCount(word string) letterCount {
	var letters = letterCount{}

	for _, r := range word {
		if unicode.IsLetter(r) {
			letters[unicode.ToUpper(r)-'A']++
		}
	}

	return letters
}
