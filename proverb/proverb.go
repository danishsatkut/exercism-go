// Package proverb provides functionality to generate "For Want of a Nail" proverb.
package proverb

import "fmt"

// Proverb generates the proverb for provided words.
func Proverb(rhyme []string) []string {
	totalWords := len(rhyme)
	lines := make([]string, 0, totalWords)

	if totalWords == 0 {
		return lines
	}

	for i := 0; i < totalWords-1; i++ {
		lines = append(lines, proverbLine(rhyme[i], rhyme[i+1]))
	}

	lines = append(lines, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return lines
}

func proverbLine(action, consequence string) string {
	return fmt.Sprintf("For want of a %s the %s was lost.", action, consequence)
}
