// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
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
