package etl

import "strings"

// Transform performs transformation for scrabble score.
func Transform(scores map[int][]string) map[string]int {
	var t = make(map[string]int)

	for score, letters := range scores {
		for _, l := range letters {
			t[strings.ToLower(l)] = score
		}
	}

	return t
}
