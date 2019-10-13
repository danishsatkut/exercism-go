package anagram

import (
	"strings"
)

// Detect returns anagrams for a word from possible candidates.
func Detect(word string, candidates []string) []string {
	var anagrams = make([]string, 0, len(candidates))

	for _, candidate := range candidates {
		if isAnagram(word, candidate) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

func isAnagram(word, candidate string) bool {
	if len(word) != len(candidate) {
		return false
	}

	word, candidate = strings.ToUpper(word), strings.ToUpper(candidate)

	if word == candidate {
		return false
	}

	s := newLetterCount(word)
	c := newLetterCount(candidate)

	return s == c
}

type letterCount [26]int

func newLetterCount(word string) letterCount {
	var letters = letterCount{}

	for _, r := range word {
		letters[r - 'A']++
	}

	return letters
}

func (l letterCount) equal(other letterCount) bool {
	for r, count := range l {
		if other[r] != count {
			return false
		}
	}

	return true
}
