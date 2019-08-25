// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	return divisibleBy(year, 4) &&
		(!divisibleBy(year, 100) || divisibleBy(year, 400))
}

func divisibleBy(year, n int) bool {
	return year%n == 0
}
