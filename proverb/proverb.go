// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var lines []string

	if len(rhyme) == 0 {
		return lines
	}

	for i := 0; i < len(rhyme) - 1; i++ {
		line := fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1])

		lines = append(lines, line)
	}

	lines = append(lines, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))

	return lines
}
