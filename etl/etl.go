package etl

import "strings"

// Transform performs transformation for scrabble score.
func Transform(scores map[int][]string) map[string]int {
	var output = make(map[string]int)

	for score, letters := range scores {
		for _, l := range letters {
			output[strings.ToLower(l)] = score
		}
	}

	return output
}
