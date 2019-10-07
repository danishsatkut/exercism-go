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

	subjectLetterCount := calculateLetterCount(word)
	candidateLetterCount := calculateLetterCount(candidate)

	for r, count := range subjectLetterCount {
		if candidateLetterCount[r] != count {
			return false
		}
	}

	return true
}

func calculateLetterCount(word string) map[rune]int {
	var letterCount = make(map[rune]int)

	for _, r := range word {
		letterCount[r]++
	}

	return letterCount
}
