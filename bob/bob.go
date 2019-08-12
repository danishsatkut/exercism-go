// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// YELL
// Question?
// YELLED QUESTION?

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	if isQuestion(remark) {
		if isNumerical(remark) {
			return "Sure."
		}

		if isYelled(remark) {
			return "Calm down, I know what I'm doing!"
		}

		return "Sure."
	}


	if isYelled(remark) {
		if isNumerical(remark) {
			return "Whatever."
		}

		return "Whoa, chill out!"
	}

	return "Whatever."
}

func isNumerical(remark string) bool {
	return strings.ToUpper(remark) == strings.ToLower(remark)
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isYelled(remark string) bool {
	return strings.ToUpper(remark) == remark
}
