// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	if strings.HasSuffix(remark, "?") {
		if strings.ToUpper(remark) == strings.ToLower(remark) {
			return "Sure."
		}

		if strings.ToUpper(remark) == remark {
			return "Calm down, I know what I'm doing!"
		}

		return "Sure."
	}


	if strings.ToUpper(remark) == remark {
		if strings.ToUpper(remark) == strings.ToLower(remark) {
			return "Whatever."
		}

		return "Whoa, chill out!"
	}

	return "Whatever."
}
