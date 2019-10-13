package anagram

import (
	"strings"
	"unicode"
)

// Detect returns anagrams for a word from possible candidates.
func Detect(word string, candidates []string) []string {
	var anagrams = make([]string, 0, len(candidates))

	sc := letterCount(word)

	for _, candidate := range candidates {
		if len(word) != len(candidate) || strings.EqualFold(word, candidate) {
			continue
		}

		if sc == letterCount(candidate) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

type counts [26]int

func letterCount(word string) counts {
	var count = counts{}

	for _, r := range word {
		if unicode.IsLetter(r) {
			count[unicode.ToUpper(r)-'A']++
		}
	}

	return count
}
