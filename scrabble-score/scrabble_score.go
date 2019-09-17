// Package scrabble provides functionality to calculate scrabble score
package scrabble

import "strings"

var alphabetScores = map[int]string{
	1:  "AEIOULNRST",
	2:  "DG",
	3:  "BCMP",
	4:  "FHVWY",
	5:  "K",
	8:  "JX",
	10: "QZ",
}

// Score calculates scrabble score for a word
func Score(word string) int {
	var score int
	letters := []rune(strings.ToUpper(word))

	for _, l := range letters {
		score += letterScore(l)
	}

	return score
}

func letterScore(letter rune) int {
	for score, alphabets := range alphabetScores {
		if strings.ContainsRune(alphabets, letter) {
			return score
		}
	}

	panic("not a valid alphabet: " + string(letter))
}
