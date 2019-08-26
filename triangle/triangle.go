// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if !isTriangle(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if (a != b && b == c) || (a != b && a == c) || (a != c && a == b) {
		return Iso
	}

	if a != b && a != c && b != c {
		return Sca
	}

	return k
}

func isSideValid(s float64) bool {
	return !(s <= 0 || math.IsNaN(s) || math.IsInf(s, 0))
}

func isTriangleInEqual(a, b, c float64) bool {
	return isInEqual(a, b, c) && isInEqual(a, c, b) && isInEqual(b, c, a)
}

func isInEqual(x, y, z float64) bool {
	return z <= x+y
}

func isTriangle(a, b, c float64) bool {
	return isSideValid(a) && isSideValid(b) && isSideValid(c) && isTriangleInEqual(a, b, c)
}
