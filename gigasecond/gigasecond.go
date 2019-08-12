// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

import (
	"math"
	"time"
)

var gigasecond = time.Duration(math.Pow10(9)) * time.Second

// AddGigasecond adds 10 ^ 9 seconds to specified time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
