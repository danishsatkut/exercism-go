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
	t := triangle{a, b, c}

	return t.kind()
}

type triangle struct {
	a, b, c float64
}

func (t triangle) isInEqual() bool {
	return isInEqual(t.a, t.b, t.c) && isInEqual(t.a, t.c, t.b) && isInEqual(t.b, t.c, t.a)
}

func (t triangle) isValid() bool {
	return isSideValid(t.a) && isSideValid(t.b) && isSideValid(t.c) && t.isInEqual()
}

func (t triangle) kind() Kind {
	if !t.isValid() {
		return NaT
	}

	if t.a == t.b && t.b == t.c {
		return Equ
	}

	if (t.b == t.c) || (t.a == t.c) || (t.a == t.b) {
		return Iso
	}

	if t.a != t.b && t.a != t.c && t.b != t.c {
		return Sca
	}

	return NaT
}

func isSideValid(s float64) bool {
	return !(s <= 0 || math.IsNaN(s) || math.IsInf(s, 0))
}

func isInEqual(x, y, z float64) bool {
	return z <= x+y
}
