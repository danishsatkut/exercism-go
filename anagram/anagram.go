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

	subjectLetterCount := newLetterCount(word)
	candidateLetterCount := newLetterCount(candidate)

	return subjectLetterCount.equal(candidateLetterCount)
}

type letterCount map[rune]int

func newLetterCount(word string) letterCount {
	var letterCount = make(map[rune]int)

	for _, r := range word {
		letterCount[r]++
	}

	return letterCount
}

func (l letterCount) equal(other letterCount) bool {
	for r, count := range l {
		if other[r] != count {
			return false
		}
	}

	return true
}
